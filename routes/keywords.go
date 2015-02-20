package routes

import (
	"fmt"
	"xavier/app"
)

func UserQuestionsKeywordsIndex(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("question")
	k, err := c.KeywordStorage.All(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, "Keywords could not be found"}
	}
	return c.JSON(200, "keywords", k)
}

func UserQuestionsKeywordsShow(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("question")
	// if q, err := c.QuestionStorage.Find(uuid, c.GetUserID()); err == nil {
	// 	// if k, err := c.KeywordStorage.All(q.ID, c.GetUserID()); err == nil {
	// 	// 	q.Keywords = k
	// 	// }
	// 	return c.JSON(200, "questions", q)
	// }
	return &app.Error{404, fmt.Sprintf("Question could not be found for uuid %s", uuid)}
}
