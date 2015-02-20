package postgres

import (
	"xavier/storage"

	"github.com/jmoiron/sqlx"
)

type NoteDatabase struct {
	*sqlx.DB
}

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

func (db *NoteDatabase) Insert(note *storage.Note) error {
	return db.Get(note, "INSERT INTO notes (content, created_at, updated_at, user_id) VALUES ($1, $2, $3, $4) RETURNING uuid, content, created_at, updated_at, user_id", note.Content, note.CreatedAt, note.UpdatedAt, note.UserID)
}
