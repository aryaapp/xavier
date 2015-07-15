package api

import (
	"fmt"
	"net/http"

	"github.com/aryaapp/xavier/storage"
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

func (a *AppContext) NewNote(c *echo.Context) error {
	var input storage.NoteInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(422, Error{"Invalid input given."})
	}

	n, err := a.NoteStorage.New(&input, c.Get("user.id").(int), 1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Note could not be created."})
	}
	return c.JSON(http.StatusCreated, Data{n})
}
