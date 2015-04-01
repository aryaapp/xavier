package storage

import (
	"net/http"
	"time"
)

type AccessToken struct {
	UUID         string
	Expiration   time.Time
	SignedString string
}

type AccessTokenJSON struct {
	Bearer string `json:"bearer"`
}

type AccessTokenStorage interface {
	New(userID string) (*AccessToken, error)
	FromRequest(r *http.Request) (*AccessToken, error)
	Parse(s string) (*AccessToken, error)
}
