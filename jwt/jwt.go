package jwt

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type JWT interface {
	Claims() map[string]interface{}
	Decode(map[string]interface{}) error
}

type Client struct {
	PrivateKey string
}

func (c *Client) Decode(obj JWT, signedString string) error {
	jwt, err := c.parseString(signedString)
	if err != nil {
		return err
	}
	return obj.Decode(jwt.Claims)
}

func (c *Client) DecodeFromRequest(obj JWT, request *http.Request) error {
	jwt, err := c.parseFromRequest(request)
	if err != nil {
		return err
	}
	return obj.Decode(jwt.Claims)
}

func (c *Client) Encode(obj JWT) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = obj.Claims()
	return token.SignedString([]byte(c.PrivateKey))
}

// Private

func (c *Client) parseString(signedString string) (*jwt.Token, error) {
	jwt, err := jwt.Parse(signedString, c.keyFunc())
	if err != nil {
		return nil, err
	}
	if !jwt.Valid {
		return nil, errors.New("token is not valid")
	}
	return jwt, nil
}

func (c *Client) parseFromRequest(request *http.Request) (*jwt.Token, error) {
	jwt, err := jwt.ParseFromRequest(request, c.keyFunc())
	if err != nil {
		return nil, err
	}
	if !jwt.Valid {
		return nil, errors.New("token is not valid")
	}
	return jwt, nil
}

func (c *Client) keyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(c.PrivateKey), nil
	}
}
