package routes

import (
	"fmt"
	"xavier/api"
	"xavier/storage"
)

type NoteParams struct {
	Content string `json:"content" validate:"nonzero"`
}

func UserNotesIndex(c *api.Context) *api.Error {
	n, err := c.NoteStorage.All(c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &api.Error{404, "Notes could not be found."}
	}
	return c.JSON(200, "notes", n)
}

func UserNotesShow(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("note")
	n, err := c.NoteStorage.Find(uuid, c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &api.Error{404, fmt.Sprintf("Note could not be found for uuid %s", uuid)}
	}
	return c.JSON(200, "notes", n)
}

func UserNotesCreate(c *api.Context) *api.Error {
	var params NoteParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		c.LogError(err)
		return &api.Error{422, "Note could not be created. Invalid parameters"}
	}

	note, err := c.NoteStorage.Insert(&storage.NoteEntry{
		Content: params.Content,
		UserID:  c.GetUserID(),
	})
	if err != nil {
		c.LogError(err)
		return &api.Error{500, "Note could not be created." + err.Error()}
	}
	return c.JSON(201, "notes", note)

}
