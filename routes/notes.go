package routes

import (
	"fmt"
	"time"
	"xavier/app"
	"xavier/storage"
)

type NoteParams struct {
	Content string `json:"content" validate:"nonzero"`
}

func UserNotesIndex(c *app.Context) *app.Error {
	n, err := c.NoteStorage.All(c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, "Notes could not be found."}
	}
	return c.JSON(200, "notes", n)
}

func UserNotesShow(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("note")
	n, err := c.NoteStorage.Find(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, fmt.Sprintf("Note could not be found for uuid %s", uuid)}
	}
	return c.JSON(200, "notes", n)
}

func UserNotesCreate(c *app.Context) *app.Error {
	var params NoteParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		c.LogError(err)
		return &app.Error{422, "Note could not be created. Invalid parameters"}
	}

	note := &storage.Note{}
	note.Content = params.Content
	note.CreatedAt = time.Now()
	note.UpdatedAt = note.CreatedAt
	note.UserID = c.GetUserID()

	if err := c.NoteStorage.Insert(note); err != nil {
		c.LogError(err)
		return &app.Error{500, "Note could not be created." + err.Error()}
	}
	return c.JSON(201, "notes", note)

}
