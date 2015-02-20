package httpSignature

import (
	"encoding/base64"
	"errors"
	"strings"
)

type Parameters struct {
	KeyId     string
	Algorithm string
	Headers   []string
}

type Signature struct {
	Scheme          string
	Parameters      Parameters
	SignedSignature []byte
}

func SignatureFromHeader(header string) (*Signature, error) {
	params := parseHeaderString(header)

	keyId, ok := params["keyId"]
	if !ok || len(keyId) == 0 {
		return nil, errors.New("Signature does not contain parameter 'keyId'.")
	}

	algorithm, ok := params["algorithm"]
	if !ok || len(algorithm) == 0 {
		return nil, errors.New("Signature does not contain parameter 'algorithm'.")
	}

	encodedSignature, ok := params["signature"]
	if !ok || len(encodedSignature) == 0 {
		return nil, errors.New("Signature does not contain parameter 'signature'.")
	}

	headersString, ok := params["headers"]
	if !ok || len(headersString) == 0 {
		return nil, errors.New("Signature does not contain parameter 'headers'.")
	}

	headers := strings.Split(headersString, " ")
	if len(headers) == 0 {
		return nil, errors.New("The 'headers' parameter does not contain any values.")
	}

	for _, header := range headers {
		if header == "signature" {
			return nil, errors.New("The 'headers' parameter is invalid.")
		}
	}

	signature, err := base64.StdEncoding.DecodeString(encodedSignature)
	if err != nil {
		return nil, errors.New("The 'signature' parameter does not contain a valid base64 encoded string.")
	}

	p := Parameters{KeyId: keyId, Algorithm: algorithm, Headers: headers}
	return &Signature{Parameters: p, SignedSignature: []byte(signature)}, nil
}

func parseHeaderString(query string) map[string]string {
	m := make(map[string]string)
	for query != "" {
		key := query
		if i := strings.IndexAny(key, ",;"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}

		if strings.Count(value, "\"") == 2 {
			m[key] = strings.Replace(value, "\"", "", 2)
		}
	}
	return m
}
