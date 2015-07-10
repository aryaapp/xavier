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

type NoteInput struct {
	Content string `json:"content" validate:"nonzero"`
}

type NoteStorage interface {
	FindAll(userID int) ([]Note, error)
	FindByUUID(uuid string, userID int) (*Note, error)
	New(input *NoteInput, userID int, appID int) (*Note, error)
}
