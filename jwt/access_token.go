package jwt

import (
	"errors"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type AccessToken struct {
	ID         string
	Expiration time.Time
	UserID     int
}

func NewAccessToken(userID int) *AccessToken {
	return &AccessToken{
		ID:         uuid.New(),
		Expiration: time.Now().Add(time.Hour * 24 * 7),
		UserID:     userID,
	}
}

func (a *AccessToken) Claims() map[string]interface{} {
	return map[string]interface{}{
		"exp": a.Expiration.Unix(),
		"jti": a.ID,
		"sub": a.UserID,
	}
}

func (a *AccessToken) Decode(claims map[string]interface{}) error {
	jti, ok := claims["jti"].(string)
	if !ok {
		return errors.New("couldn't extract token id")
	}
	sub, ok := claims["sub"].(float64)
	if !ok {
		return errors.New("couldn't extract user id")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("found no expiration date in token")
	}

	a.ID = jti
	a.Expiration = time.Unix(int64(exp), 0).In(time.UTC)
	a.UserID = int(sub)
	return nil
}
