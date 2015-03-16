package postgres

import (
	"time"
	"xavier/storage"

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

func (db *NoteDatabase) All(userID int) ([]storage.Note, error) {
	n := []storage.Note{}
	err := db.Select(&n, "SELECT uuid, content, created_at, updated_at, user_id FROM notes n WHERE n.user_id = $1", userID)
	return n, err
}

func (db *NoteDatabase) Find(uuid string, userID int) (*storage.Note, error) {
	n := &storage.Note{}
	err := db.Get(n, "SELECT uuid, content, created_at, updated_at, user_id FROM notes n WHERE n.uuid = $1 AND n.user_id = $2 LIMIT 1", uuid, userID)
	return n, err
}

func (db *NoteDatabase) Insert(noteEntry *storage.NoteEntry) (*storage.Note, error) {
	n := &storage.Note{}
	createdAt := time.Now()
	return n, db.Get(n, notesInsert, noteEntry.Content, createdAt, createdAt, noteEntry.UserID)
}
