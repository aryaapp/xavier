package httpSignature

import (
	"errors"
	"net/http"

	"crypto/hmac"
	"crypto/sha256"
	//"encoding/base64"
)

const newline = "\u000A"
const requestTargetString = "(request-target)"

type SigningOptions struct {
	Headers       map[string]string
	RequestTarget string
	SecretKey     string
}

func (s *Signature) SignWithRequest(request *http.Request, secretKey string) error {

	options := SigningOptions{SecretKey: secretKey}

	return s.SignWithOptions(options)
}

func (s *Signature) SignWithOptions(options SigningOptions) error {
	bs := []byte{}
	for _, key := range s.Parameters.Headers {
		if key == requestTargetString {
			bs = append(bs, (key + ":" + options.RequestTarget + newline)...)
			continue
		}

		value, ok := options.Headers[key]
		if !ok {
			return errors.New("Request does not contain header.")
		}

		bs = append(bs, (key + ":" + value + newline)...)
	}

	hmac := hmac.New(sha256.New, []byte(options.SecretKey))
	hmac.Write(bs)
	s.SignedSignature = hmac.Sum(nil)
	return nil
}
