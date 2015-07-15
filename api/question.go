package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppContext) FindQuestionByUUID(c *echo.Context) error {
	uuid, userID := c.Param("uuid"), c.Get("user.id").(int)
	q, err := a.QuestionStorage.FindByUUID(uuid, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Question could not be found."})
	}
	return c.JSON(http.StatusOK, Data{q})
}
