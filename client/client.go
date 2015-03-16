package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

const (
	defaultBaseURL   = "http://localhost:8000"
	defaultMediaType = "application/json; charset=utf-8"
)

type Client struct {
	AppID         string
	AppSecret     string
	TokenProvider TokenProvider
	client        *http.Client

	Themes ThemeService
}

type ClientError struct {
	Code    int
	Message string
}

type ClientKeyfunc func(*http.Response) (string, interface{})

type TokenProvider interface {
	GetAccessToken() string
	SetAccessToken(accessToken string) error
}

func NewClient(appID, appSecret string, tokenProvider TokenProvider) *Client {
	client := &Client{
		AppID:         appID,
		AppSecret:     appSecret,
		TokenProvider: tokenProvider,
		client:        http.DefaultClient,
	}

	client.Themes = ThemeService{client}
	return client
}

func (c *Client) NewRequest(method, verb string, tokenRequired bool, body interface{}) (*http.Request, error) {
	buf, err := c.parseBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.parseURL(verb), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", defaultMediaType)
	req.Header.Add("Content-Type", defaultMediaType)

	if tokenRequired {
		req.Header.Add("Authorization", "Bearer "+c.TokenProvider.GetAccessToken())
	} else {
		req.Header.Add("Authorization", "App "+c.AppID)
	}
	return req, nil
}

func (c *Client) Do(request *http.Request, key string, value interface{}) (*http.Response, error) {
	response, err := c.client.Do(request)

	log.Println("piss 0")

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	log.Println("yolo 0")

	responseJSON := map[string]*json.RawMessage{}
	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return response, err
	}

	log.Println("yolo 1")

	responseData, ok := responseJSON[key]
	if !ok {
		return response, errors.New("Invalid key")
	}
	log.Println("yolo 2")

	if err := json.Unmarshal(*responseData, value); err != nil {
		return response, err
	}

	// if errObj, ok := test["error"]; ok {
	// 	clientErr := ClientError{}
	// 	err = json.Unmarshal(errObj, &clientErr)
	// 	return clientErr, response, err
	// }

	// themes := []storage.Theme{}
	// json.Unmarshal(test["themes"], &themes)

	// log.Println(key, value)

	// test := map[string]interface{}{}
	// if v != nil {
	// 	if w, ok := v.(io.Writer); ok {
	// 		io.Copy(w, response.Body)
	// 	} else {
	// err = json.NewDecoder(response.Body).Decode(&test)

	log.Println("yolo 3")
	log.Println(value)
	// log.Println(themes)
	// 	}
	// }
	return response, nil
}

// type ClientResponse struct {
// 	*http.Response
// }

// func (cr *ClientResponse) Value(key string, value interface{}) error {

// 	log.Println(cr.Body)

// 	if err := json.NewDecoder(cr.Body).Decode(&value); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (cr *ClientResponse) Error(key string, value *app.Error) error {
// 	if err := json.NewDecoder(cr.Body).Decode(&value); err != nil {
// 		return err
// 	}
// 	return nil
// }

// Private

func (c *Client) parseURL(urlStr string) string {
	return c.sanitize(defaultBaseURL) + "/" + c.sanitize(urlStr)
}

func (c *Client) sanitize(path string) string {
	last := len(path) - 1
	if last >= 0 && path[last] == '/' {
		path = path[:last]
		last--
	}
	if last >= 0 && path[0] == '/' {
		path = path[1:]
	}
	return path
}

func (c *Client) parseBody(body interface{}) (io.ReadWriter, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}
