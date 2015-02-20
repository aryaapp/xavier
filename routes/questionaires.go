package routes

import (
	"fmt"
	"xavier/app"
)

func UserQuestionairesIndex(c *app.Context) *app.Error {
	if q, err := c.QuestionaireStorage.All(c.GetUserID()); err == nil {
		return c.JSON(200, "questionaires", q)
	}
	return &app.Error{404, "Questionaires could not be found."}
}

func UserQuestionairesShow(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("questionaire")
	if q, err := c.QuestionaireStorage.Find(uuid, c.GetUserID()); err == nil {
		return c.JSON(200, "questionaires", q)
	}
	return &app.Error{404, fmt.Sprintf("Questionaire could not be found for uuid %s", uuid)}
}
