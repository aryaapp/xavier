package app

import "net/http"

func Handler(ctx *Context, h func(*Context) *Error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := ChildContext(ctx, r, w)
		if e := h(c); e != nil {
			switch e.Code {
			case 422:
				c.JSONError(e.Code, "Unprocessable Entity", e.Message)
			default:
				c.JSONError(e.Code, http.StatusText(e.Code), e.Message)
			}
		}
	}
}
