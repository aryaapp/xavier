package storage

import "time"

type Note struct {
	ID        int       `json:"-" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	AppID     int       `json:"-" db:"app_id"`
	UserID    int       `json:"-" db:"user_id"`
}

type NoteEntry struct {
	Content string `json:"content"`
	AppID   int    `json:"app_id"`
	UserID  int    `json:"user_id"`
}

type NoteStorage interface {
	All(int) ([]Note, error)
	Find(string, int) (*Note, error)
	Insert(*NoteEntry) (*Note, error)
}
