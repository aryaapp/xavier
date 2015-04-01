package jwt

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

///// Signer

type Signer struct {
	signKey []byte
}

func NewSigner(privKey []byte) (*Signer, error) {
	return &Signer{privKey}, nil
}

func (s *Signer) Sign(claims map[string]interface{}) (string, error) {
	t := jwt.New(jwt.SigningMethodRS256)
	t.Claims = claims
	return t.SignedString(s.signKey)
}

///// Verifier

type Verifier struct {
	verifyKey []byte
}

// GetKey can be passed to the `jwt.Parse` function and returns the key which is
// used for decoding.
func (v *Verifier) GetKey(token *jwt.Token) (interface{}, error) {
	return v.verifyKey, nil
}

func NewVerifier(pubKey []byte) (*Verifier, error) {
	return &Verifier{pubKey}, nil
}

func (v *Verifier) parseString(s string) (*jwt.Token, error) {
	jwt, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return v.verifyKey, nil
	})
	if err != nil || !jwt.Valid {
		return nil, errors.New("Invalid Signup Token")
	}
	return jwt, nil
}

func (v *Verifier) parseFromRequest(r *http.Request) (*jwt.Token, error) {
	jwt, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return v.verifyKey, nil
	})
	if err != nil || !jwt.Valid {
		return nil, errors.New("Invalid Token")
	}
	return jwt, nil
}
