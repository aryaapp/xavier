package api

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

type middleware struct {
	ctx        *Context
	handleFunc func(*Context, http.HandlerFunc) *Error
}

func (m middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c := ChildContext(m.ctx, r, w)
	if e := m.handleFunc(c, next); e != nil {
		switch e.Code {
		case 422:
			c.JSONError(e.Code, "Unprocessable Entity", e.Message)
		default:
			c.JSONError(e.Code, http.StatusText(e.Code), e.Message)
		}
	}
}

func Middleware(ctx *Context, handleFunc func(*Context, http.HandlerFunc) *Error) negroni.Handler {
	return middleware{ctx, handleFunc}
}
