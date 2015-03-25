package routes

import (
	"fmt"
	"xavier/api"
)

func UserQuestionsShow(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("question")
	q, err := c.QuestionStorage.Find(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &api.Error{404, fmt.Sprintf("Question could not be found for uuid %s", uuid)}
	}
	return c.JSON(200, "questions", q)
}
