package routes

import (
	"fmt"
	"xavier/api"
)

func UserQuestionsKeywordsIndex(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("question")
	k, err := c.KeywordStorage.All(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &api.Error{404, "Keywords could not be found"}
	}
	return c.JSON(200, "keywords", k)
}

func UserQuestionsKeywordsShow(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("question")
	// if q, err := c.QuestionStorage.Find(uuid, c.GetUserID()); err == nil {
	// 	// if k, err := c.KeywordStorage.All(q.ID, c.GetUserID()); err == nil {
	// 	// 	q.Keywords = k
	// 	// }
	// 	return c.JSON(200, "questions", q)
	// }
	return &api.Error{404, fmt.Sprintf("Question could not be found for uuid %s", uuid)}
}
