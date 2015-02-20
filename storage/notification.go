package storage

import "time"

type Notification struct {
	ID         int       `json:"-" db:"id"`
	UUID       string    `json:"uuid" db:"uuid"`
	Message    string    `json:"message" db:"message"`
	Read       bool      `json:"read" db:"read"`
	ObjectID   int       `json:"object_id" db:"object_id"`
	ObjectType string    `json:"object_type" db:"object_type"`
	ObjectUri  string    `json:"object_uri" db:"object_uri"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	UserID     int       `json:"-" db:"user_id"`
}

type NotificationStorage interface {
	Find(string, int) (*Notification, error)
	Insert(*Notification) error
}
