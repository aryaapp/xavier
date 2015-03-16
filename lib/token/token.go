package token

import (
	"errors"
	"net/http"
	"strings"
	"time"
	s "xavier/lib/util/strings"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Audience string /* client */
	Expires  time.Time
	JTI      string
	Subject  int /* user */
	Scopes   []string
}

var (
	ErrTokenEmpty           = errors.New("Token could not be found.")
	ErrTokenParse           = errors.New("Token could not be parsed.")
	ErrTokenExpired         = errors.New("Token is expired.")
	ErrTokenInvalidAudience = errors.New("Token contains an invalid audience.")
	ErrTokenInvalidSubject  = errors.New("Token contains an invalid subject.")
	ErrTokenInvalidScopes   = errors.New("Token should contain at least one scope.")
)

func New(audience string, expiration time.Duration, subject int, scopes []string) *Token {
	return &Token{
		Audience: audience,
		Expires:  time.Now().Add(expiration),
		JTI:      s.Random(10),
		Scopes:   scopes,
		Subject:  subject,
	}
}

func ParseFromRequest(r *http.Request, secret string) (*Token, error) {
	t := &Token{}
	j, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return t, ErrTokenEmpty
	}

	if err != nil {
		if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorExpired != 0 {
			return t, ErrTokenExpired
		}
		return t, ErrTokenParse
	}
	if aud, ok := j.Claims["aud"].(string); ok {
		t.Audience = aud
		if t.Audience == "" {
			return t, ErrTokenInvalidAudience
		}
	}
	if exp, ok := j.Claims["exp"].(float64); ok {
		t.Expires = time.Unix(int64(exp), 0)
	}
	if jti, ok := j.Claims["jti"].(string); ok {
		t.JTI = jti
	}
	if scp, ok := j.Claims["scp"].(string); ok {
		t.Scopes = strings.Split(scp, ",")
		if len(t.Scopes) == 0 {
			return t, ErrTokenInvalidScopes
		}
	}
	if sub, ok := j.Claims["sub"].(float64); ok {
		t.Subject = int(sub)
		if t.Subject == 0 {
			return t, ErrTokenInvalidSubject
		}
	}
	return t, nil
}

func (t *Token) Sign(secret string) (string, error) {
	j := jwt.New(jwt.GetSigningMethod("HS256"))
	j.Claims["aud"] = t.Audience
	j.Claims["exp"] = int64(t.Expires.Unix())
	j.Claims["jti"] = t.JTI
	j.Claims["scp"] = t.Scopes
	j.Claims["sub"] = t.Subject
	return j.SignedString([]byte(secret))
}

func (t *Token) ContainsScopes(scopes []string) bool {
	return s.ContainsSlice(t.Scopes, scopes)
}

func scopesToString(scp []string) string {
	var bs []byte
	for i, v := range scp {
		bs = append(bs, v...)
		if i < (len(scp) - 1) {
			bs = append(bs, ","...)
		}
	}
	return string(bs)
}
