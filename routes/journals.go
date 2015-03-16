package routes

import (
	"encoding/json"
	"fmt"
	"time"
	"xavier/app"
	"xavier/lib/util/pg"
	"xavier/storage"
)

type AnswerParams struct {
	QuestionID string      `json:"question_id" validate:"nonzero,uuid"`
	Answer     interface{} `json:"answer"`
}

type JournalParams struct {
	Answers  []AnswerParams         `json:"answers"`
	Metadata map[string]interface{} `json:"metadata"`
}

func UserJournalsIndex(c *app.Context) *app.Error {
	j, err := c.JournalStorage.All(c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, "Journals could not be found."}
	}
	return c.JSON(200, "journals", j)
}

func UserJournalsShow(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("journal")
	j, err := c.JournalStorage.Find(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, fmt.Sprintf("Journal could not be found for uuid %s", uuid)}
	}

	a, err := c.JournalStorage.Answers(j.ID)
	if err != nil {
		c.LogError(err)
	}
	j.Answers = a

	return c.JSON(200, "journals", j)
}

func UserJournalsCreate(c *app.Context) *app.Error {
	var params JournalParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		c.LogError(err)
		return &app.Error{422, "Journal could not be created. Invalid parameters: " + err.Error()}
	}

	if len(params.Answers) == 0 {
		return &app.Error{422, "Journal could not be created. No answers are provided."}
	}

	if err := c.Validator.Validate(params.Answers); err != nil {
		c.LogError(err)
		return &app.Error{422, "Journal could not be created. Invalid answers: " + err.Error()}
	}

	questionsIDs := make([]string, 0, len(params.Answers))
	answersForQuestions := make(map[string]pg.JSON)
	for _, answerParam := range params.Answers {
		answerBytes, err := json.Marshal(answerParam.Answer)
		if err != nil {
			c.LogError(err)
			return &app.Error{422, fmt.Sprintf("Journal could not be created. The answer for question %s contains invalid JSON.", answerParam.QuestionID)}
		}

		questionsIDs = append(questionsIDs, answerParam.QuestionID)
		answersForQuestions[answerParam.QuestionID] = pg.JSON(string(answerBytes))
	}

	questionsBytes, err := json.Marshal(questionsIDs)
	if err != nil {
		c.LogError(err)
		return &app.Error{422, "Journal could not be created. Question IDs could not be marshalled."}
	}

	questions, _ := c.QuestionStorage.WhereIn(questionsIDs)
	if len(questions) == 0 || len(questionsIDs) != len(questions) {
		return &app.Error{422, "Journal could not be created. Invalid questions provided."}
	}

	journal := &storage.Journal{}
	journal.Feeling = -1
	journal.Questions = pg.JSON(string(questionsBytes))
	journal.CreatedAt = time.Now()
	journal.UpdatedAt = journal.CreatedAt
	journal.AppID = c.GetAppForCurrentRequest().ID
	journal.UserID = c.GetUserID()
	journal.Answers = []storage.Answer{}

	for _, question := range questions {
		answer := storage.Answer{}
		answer.Values = answersForQuestions[question.UUID]
		answer.Answered = string(answer.Values) != "null"
		answer.QuestionID = question.ID
		journal.Answers = append(journal.Answers, answer)
		if question.Processor == "emotions" {
			feelings := map[string]interface{}{}
			answer.Values.Unmarshal(&feelings)

			feeling, ok := feelings["feeling"].(float64)
			if !ok {
				return &app.Error{422, "Journal could not be created. Invalid feeling provided."}
			}
			journal.Feeling = int(feeling)
		}
	}

	if err := c.JournalStorage.Insert(journal); err != nil {
		c.LogError(err)
		return &app.Error{500, "Journal could not be created." + err.Error()}
	}

	go func(c *app.Context, journal *storage.Journal, questions []storage.Question) {
		for i, question := range questions {
			if question.Autocompletes {
				names := []string{}
				answer := journal.Answers[i]
				answer.Values.Unmarshal(&names)

				for _, name := range names {
					keyword := &storage.Keyword{}
					keyword.Name = name
					keyword.QuestionID = question.ID
					keyword.UserID = journal.UserID
					if err := c.KeywordStorage.InsertOrUpdate(keyword); err != nil {
						c.LogError(err)
					}
				}
			}
		}
	}(c, journal, questions)

	return c.JSON(201, "journals", journal)
}
