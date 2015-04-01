package jwt

import (
	"errors"
	"net/http"
	"time"

	"xavier/storage"
)

type AccessTokenStorage struct {
	*Signer
	*Verifier
}

// New returns a new AccessToken which can be used for encoding to JWT.
func (a *AccessTokenStorage) New(uuid string) (*storage.AccessToken, error) {
	token := &storage.AccessToken{
		UUID:       uuid,
		Expiration: time.Now().Add(time.Hour * 24 * 7),
	}

	signedString, err := a.Sign(map[string]interface{}{
		"exp": token.Expiration.Unix(),
		"sub": token.UUID,
	})
	if err != nil {
		return nil, err
	}

	token.SignedString = signedString
	return token, nil
}

// FromRequest parses an bearer string found in the request to an access token.
func (a *AccessTokenStorage) FromRequest(r *http.Request) (*storage.AccessToken, error) {
	token, err := a.parseFromRequest(r)
	if err != nil {
		return nil, errors.New("invalid Access Token")
	}

	// TODO: exact same functionality in `Parse` function below, should be
	// extracted.
	uuid, ok := token.Claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid Access Token claims")
	}

	exp, ok := token.Claims["exp"].(float64)
	if !ok {
		return nil, errors.New("couldn't extract expiration date from token")
	}

	expTime := time.Unix(int64(exp), 0)
	at := &storage.AccessToken{
		UUID:       uuid,
		Expiration: expTime,
	}
	return at, nil
}

// Parse transforms the JWT string to an access token.
func (a *AccessTokenStorage) Parse(s string) (*storage.AccessToken, error) {
	jwt, err := a.parseString(s)
	if err != nil {
		return nil, err
	}

	uuid, ok := jwt.Claims["sub"].(string)
	if !ok {
		return nil, errors.New("couldn't extract uuid from access token")
	}
	exp, ok := jwt.Claims["exp"].(float64)
	if !ok {
		return nil, errors.New("couldn't extract expiration date from token")
	}

	expTime := time.Unix(int64(exp), 0)
	at := &storage.AccessToken{
		UUID:       uuid,
		Expiration: expTime,
	}
	return at, nil
}
