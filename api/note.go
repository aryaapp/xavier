package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppContext) FindAllNotes(c *echo.Context) error {
	userID := c.Get("user.id").(int)
	n, err := a.NoteStorage.FindAll(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Notes could not be found."})
	}
	return c.JSON(http.StatusOK, Data{n})
}

func (a *AppContext) FindNoteByUUID(c *echo.Context) error {
	uuid, userID := c.Param("uuid"), c.Get("user.id").(int)
	n, err := a.NoteStorage.FindByUUID(uuid, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Note could not be found for id %s", uuid)})
	}
	return c.JSON(http.StatusOK, Data{n})
}
