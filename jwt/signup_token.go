package jwt

import (
	"errors"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type SignupToken struct {
	ID         string
	Expiration time.Time
	Email      string
	Password   string
}

func NewSignupToken(email, password string) *SignupToken {
	return &SignupToken{
		ID:         uuid.New(),
		Expiration: time.Now().Add(time.Hour * 24 * 7),
		Email:      email,
		Password:   password,
	}
}

func (s *SignupToken) Claims() map[string]interface{} {
	return map[string]interface{}{
		"exp": s.Expiration.Unix(),
		"jti": s.ID,
	}
}

func (s *SignupToken) Decode(claims map[string]interface{}) error {
	jti, ok := claims["jti"].(string)
	if !ok {
		return errors.New("couldn't extract token id")
	}
	// uid, ok := claims["sub"].(string)
	// if !ok {
	// 	return errors.New("couldn't extract email")
	// }
	exp, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("found no expiration date in token")
	}

	s.ID = jti
	s.Expiration = time.Unix(int64(exp), 0).In(time.UTC)
	return nil
}
