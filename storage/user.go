package storage

import (
	"errors"
	"time"

	"github.com/guregu/null"
)

type User struct {
	ID            int            `json:"-" db:"id"`
	UUID          string         `json:"uuid" db:"uuid"`
	Email         string         `json:"email" db:"email"`
	Password      string         `json:"-" db:"password_digest"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
	FullName      null.String    `json:"fullname" db:"fullname"`
	Gender        string         `json:"gender" db:"gender"`
	Public        bool           `json:"public" db:"public"`
	Professional  bool           `json:"professional" db:"professional"`
	AppID         int            `json:"-" db:"app_id"`
	ThemeID       int            `json:"-" db:"theme_id"`
	Journals      []Journal      `json:"journals"`
	Questionaires []Questionaire `json:"questionaires"`
	Theme         *Theme         `json:"theme,omitempty"`
}

type UserRegistration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AppID    int    `json:"app_id"`
}

type UserStorage interface {
	Find(string) (*User, error)
	FindByID(int) (*User, error)
	FindByEmail(string) (*User, error)
	Insert(*UserRegistration) (*User, error)
}

var UserConflictError = errors.New("Already exists.")
