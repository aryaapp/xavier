package jwt

import (
	"errors"

	"xavier/storage"
)

type SignupTokenStorage struct {
	*Signer
	*Verifier
}

func (s *SignupTokenStorage) New(email, username, password string) (*storage.SignupToken, error) {
	token := storage.NewSignupToken(email, username, password)
	signedString, err := s.Sign(map[string]interface{}{
		"exp": token.Expiration.Unix(),
		"user": map[string]string{
			"email":    token.Email,
			"username": token.Username,
			"password": token.Password,
		},
	})
	if err != nil {
		return nil, err
	}

	token.SignedString = signedString
	return token, nil
}

func (s *SignupTokenStorage) Parse(str string) (*storage.SignupToken, error) {
	jwt, err := s.parseString(str)
	if err != nil {
		return nil, errors.New("invalid Signup Token")
	}

	user, ok := jwt.Claims["user"].(map[string]string)
	if !ok {
		return nil, errors.New("invalid Signup Token claims")
	}
	email, emailOk := user["email"]
	username, usernameOk := user["username"]
	password, passwordOk := user["password"]

	if !emailOk || !usernameOk || !passwordOk {
		return nil, errors.New("invalid Signup Token claims")
	}
	return &storage.SignupToken{Email: email, Username: username, Password: password}, nil
}
