package routes

import (
	"fmt"
	"xavier/api"
)

func UserQuestionairesIndex(c *api.Context) *api.Error {
	if q, err := c.QuestionaireStorage.All(c.GetUserID()); err == nil {
		return c.JSON(200, "questionaires", q)
	}
	return &api.Error{404, "Questionaires could not be found."}
}

func UserQuestionairesShow(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("questionaire")
	if q, err := c.QuestionaireStorage.Find(uuid, c.GetUserID()); err == nil {
		return c.JSON(200, "questionaires", q)
	}
	return &api.Error{404, fmt.Sprintf("Questionaire could not be found for uuid %s", uuid)}
}
