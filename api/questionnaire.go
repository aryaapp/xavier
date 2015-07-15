package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppContext) FindAllQuestionaires(c *echo.Context) error {
	userID := c.Get("user.id").(int)
	q, err := a.QuestionaireStorage.FindAll(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Questionaires could not be found."})
	}
	return c.JSON(http.StatusOK, Data{q})
}

func (a *AppContext) FindQuestionaireByUUID(c *echo.Context) error {
	uuid, userID := c.Param("uuid"), c.Get("user.id").(int)
	q, err := a.QuestionaireStorage.FindByUUID(uuid, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Questionaire could not be found."})
	}
	return c.JSON(http.StatusOK, Data{q})
}
