package main

import (
	"net/http"
	"strings"
	"xavier/api"
	"xavier/lib/token"
	"xavier/storage"
)

func ContentType(c *api.Context, next http.HandlerFunc) *api.Error {
	contentTypes := map[string]string{
		"POST":  "application/json",
		"PUT":   "application/json",
		"PATCH": "application/json",
	}

	if contentType, ok := contentTypes[c.Request.Method]; ok {
		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), contentType) {
			return &api.Error{415, "Please specify a supported content-type."}
		}
	}

	next(c.Writer, c.Request)
	return nil
}

func CurrentApp(c *api.Context, next http.HandlerFunc) *api.Error {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return &api.Error{401, "Authorization error: App header is empty."}
	}

	if !(len(header) > 3 && strings.ToUpper(header[0:3]) == "APP") {
		return &api.Error{401, "Authorization error: App header is malformed."}
	}

	appID := header[4:]
	a, err := setApp(c, appID)
	if err != nil {
		return &api.Error{401, "Authorization error: Invalid App request"}
	}
	c.SetAppForCurrentRequest(a)

	next(c.Writer, c.Request)
	return nil
}

func Bearer(c *api.Context, next http.HandlerFunc) *api.Error {
	token, err := token.ParseFromRequest(c.Request, c.Environment.Secret)
	if err != nil {
		return &api.Error{401, "Authorization error: " + err.Error()}
	}

	if !token.ContainsScopes(c.Scopes) {
		return &api.Error{401, "Authorization error: Doesn't have required scopes"}
	}

	a, err := setApp(c, token.Audience)
	if err != nil {
		return &api.Error{401, "Authorization error: Invalid App in token"}
	}

	c.SetAppForCurrentRequest(a)
	c.SetUserID(token.Subject)

	next(c.Writer, c.Request)
	return nil
}

func setApp(c *api.Context, uuid string) (*storage.App, error) {
	a, err := c.AppCache.Find(uuid)
	if a == nil && err == nil {
		a, err = c.AppStorage.Find(uuid)
	}
	return a, err
}
