package routes

import (
	"fmt"
	"xavier/app"
)

func UserQuestionsShow(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("question")
	q, err := c.QuestionStorage.Find(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, fmt.Sprintf("Question could not be found for uuid %s", uuid)}
	}
	return c.JSON(200, "questions", q)
}
