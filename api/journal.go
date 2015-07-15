package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppContext) FindAllJournals(c *echo.Context) error {
	userID := c.Get("user.id").(int)
	j, err := a.JournalStorage.FindAll(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Journals could not be found."})
	}
	return c.JSON(http.StatusOK, Data{j})
}

func (a *AppContext) FindJournalByUUID(c *echo.Context) error {
	uuid, userID := c.Param("uuid"), c.Get("user.id").(int)
	j, err := a.JournalStorage.FindByUUID(uuid, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Journal could not be found for id %s", uuid)})
	}
	return c.JSON(http.StatusOK, Data{j})
}

// func (a *AppContext) NewNote(c *echo.Context) error {
// 	var input storage.NoteInput
// 	if err := c.Bind(&input); err != nil {
// 		return c.JSON(422, Error{"Invalid input given."})
// 	}

// 	n, err := a.NoteStorage.New(&input, c.Get("user.id").(int), 1)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, Error{"Note could not be created."})
// 	}
// 	return c.JSON(http.StatusCreated, Data{n})
// }
