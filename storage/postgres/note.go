package postgres

import (
	"time"

	"github.com/aryaapp/xavier/storage"

	"github.com/jmoiron/sqlx"
)

type NoteDatabase struct {
	*sqlx.DB
}

const (
	notesInsert = `
		INSERT INTO notes  (content, created_at, updated_at, user_id) VALUES  ($1, $2, $3, $4) 
		RETURNING uuid, content, created_at, updated_at, user_id
	`
)

func (db *NoteDatabase) FindAll(userID int) ([]storage.Note, error) {
	n := []storage.Note{}
	err := db.Select(&n, "SELECT uuid, content, created_at, updated_at, user_id FROM notes n WHERE n.user_id = $1", userID)
	return n, err
}

func (db *NoteDatabase) FindByUUID(uuid string, userID int) (*storage.Note, error) {
	n := &storage.Note{}
	err := db.Get(n, "SELECT uuid, content, created_at, updated_at, user_id FROM notes n WHERE n.uuid = $1 AND n.user_id = $2 LIMIT 1", uuid, userID)
	return n, err
}

func (db *NoteDatabase) New(input *storage.NoteInput, userID int, appID int) (*storage.Note, error) {
	n := &storage.Note{}
	createdAt := time.Now()
	return n, db.Get(n, notesInsert, input.Content, createdAt, createdAt, userID)
}
