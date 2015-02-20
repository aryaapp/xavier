package main

import (
	"net/http"
	"strings"
	"xavier/app"
	"xavier/lib/token"
	"xavier/storage"
)

func ContentType(c *app.Context, next http.HandlerFunc) *app.Error {
	contentTypes := map[string]string{
		"POST":  "application/json",
		"PUT":   "application/json",
		"PATCH": "application/json",
	}

	if contentType, ok := contentTypes[c.Request.Method]; ok {
		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), contentType) {
			return &app.Error{415, "Please specify a supported content-type."}
		}
	}

	next(c.Writer, c.Request)
	return nil
}

func Client(c *app.Context, next http.HandlerFunc) *app.Error {
	header := c.Request.Header.Get("Client")
	if len(header) == 0 {
		return &app.Error{401, "Authorization error: Client header is empty."}
	}

	client, err := setClient(c, header)
	if err != nil {
		return &app.Error{401, "Authorization error: Invalid Client request"}
	}
	c.SetClientForCurrentRequest(client)

	next(c.Writer, c.Request)
	return nil
}

func Bearer(c *app.Context, next http.HandlerFunc) *app.Error {
	token, err := token.ParseFromRequest(c.Request, c.Environment.Secret)
	if err != nil {
		return &app.Error{401, "Authorization error: " + err.Error()}
	}

	client, err := setClient(c, token.Audience)
	if err != nil {
		return &app.Error{401, "Authorization error: Invalid Client in token"}
	}

	c.SetClientForCurrentRequest(client)
	c.SetUserID(token.Subject)

	next(c.Writer, c.Request)
	return nil
}

func setClient(c *app.Context, uuid string) (*storage.Client, error) {
	client, err := c.ClientCache.Find(uuid)
	if client == nil && err == nil {
		client, err = c.ClientStorage.Find(uuid)
	}
	return client, err
}
