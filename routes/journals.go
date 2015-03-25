package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"xavier/api"
	"xavier/lib/util/pg"
	"xavier/storage"
)

type AnswerParams struct {
	QuestionID string          `json:"question_id" valid:"required,uuid"`
	Answer     json.RawMessage `json:"answer" valid:"required"`
}

type JournalParams struct {
	Answers  []AnswerParams         `json:"answers" valid:"required"`
	Metadata map[string]interface{} `json:"metadata"`
}

func UserJournalsIndex(c *api.Context) *api.Error {
	j, err := c.JournalStorage.All(c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &api.Error{404, "Journals could not be found."}
	}
	return c.JSON(200, "journals", j)
}

func UserJournalsShow(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("journal")
	j, err := c.JournalStorage.Find(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &api.Error{404, fmt.Sprintf("Journal could not be found for uuid %s", uuid)}
	}

	a, err := c.JournalStorage.Answers(j.ID)
	if err != nil {
		c.LogError(err)
	}
	j.Answers = a

	return c.JSON(200, "journals", j)
}

func UserJournalsCreate(c *api.Context) *api.Error {
	var params JournalParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		// c.LogError(err)
		return &api.Error{422, "Journal could not be created. Invalid parameters: " + err.Error()}
	}

	if len(params.Answers) == 0 {
		return &api.Error{422, "Journal could not be created. No answers are provided."}
	}

	// if err := api.Validate(params.Answers); err != nil {
	// 	// c.LogError(err)
	// 	return &api.Error{422, "Journal could not be created. Invalid answers: " + err.Error()}
	// }

	questionsIDs := make([]string, 0, len(params.Answers))
	answersForQuestions := make(map[string]pg.JSON)
	nullAnswers := make([]string, 0, len(params.Answers))
	for _, answerParam := range params.Answers {
		answerBytes := answerParam.Answer
		if len(answerBytes) == 0 {
			answerBytes = []byte("null")
		}

		answer := pg.JSON(string(answerBytes))
		if answer.IsNull() {
			nullAnswers = append(nullAnswers, answerParam.QuestionID)
		}
		answersForQuestions[answerParam.QuestionID] = answer
		questionsIDs = append(questionsIDs, answerParam.QuestionID)
	}

	log.Println(nullAnswers)
	if len(nullAnswers) == len(questionsIDs) {
		return &api.Error{422, "Journal could not be created. All answers contain empty values."}
	}

	questionsBytes, err := json.Marshal(questionsIDs)
	if err != nil {
		c.LogError(err)
		return &api.Error{422, "Journal could not be created. Question IDs could not be marshalled."}
	}

	questions, _ := c.QuestionStorage.WhereIn(questionsIDs)
	if len(questions) == 0 || len(questionsIDs) != len(questions) {
		return &api.Error{422, "Journal could not be created. Invalid questions provided."}
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
		answer.Answered = !answer.Values.IsNull()
		answer.QuestionID = question.ID
		journal.Answers = append(journal.Answers, answer)
		if question.Processor == "emotions" {
			feelings := map[string]interface{}{}
			answer.Values.Unmarshal(&feelings)

			feeling, ok := feelings["feeling"].(float64)
			if !ok {
				return &api.Error{422, "Journal could not be created. Invalid feeling provided."}
			}
			journal.Feeling = int(feeling)
		}
	}

	if err := c.JournalStorage.Insert(journal); err != nil {
		c.LogError(err)
		return &api.Error{500, "Journal could not be created." + err.Error()}
	}

	go func(c *api.Context, journal *storage.Journal, questions []storage.Question) {
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
