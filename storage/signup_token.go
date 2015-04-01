package storage

import "time"

type SignupToken struct {
	Email        string
	Username     string
	Password     string
	Expiration   time.Time
	SignedString string
}

type SignupTokenJSON struct {
	Message string `json:"message"`
}

type SignupTokenStorage interface {
	New(string, string, string) (*SignupToken, error)
	Parse(string) (*SignupToken, error)
}

func NewSignupToken(email string, username string, password string) *SignupToken {
	return &SignupToken{
		Email:      email,
		Username:   username,
		Password:   password,
		Expiration: time.Now().Add(time.Hour * 24 * 2),
	}
}
