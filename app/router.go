package app

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/nbio/httpcontext"
)

type Router struct {
	router *httprouter.Router
	groups []group
}

func NewRouter() *Router {
	return &Router{router: httprouter.New()}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.Handle("GET", path, handler)
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.Handle("PUT", path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.Handle("POST", path, handler)
}

func (r *Router) Patch(path string, handler http.HandlerFunc) {
	r.Handle("PATCH", path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.Handle("DELETE", path, handler)
}

func (r *Router) Options(path string, handler http.HandlerFunc) {
	r.Handle("OPTIONS", path, handler)
}

func (r *Router) Handle(method string, path string, handler http.HandlerFunc) {
	path = r.sanitize(path)
	if len(r.groups) > 0 {
		ln := len(r.groups)
		for i := range r.groups {
			g := r.groups[ln-1-i]
			n := negroni.New(g.middlewares...)
			n.UseHandler(handler)
			handler = n.ServeHTTP

			if len(path) > 0 {
				path = "/" + path
			}
			if len(g.path) > 0 {
				path = "/" + g.path + path
			}
		}
	}

	r.router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		setParams(r, p)
		handler(w, r)
	})
}

//////////////////////////////////////////////////////////////////////

type group struct {
	path        string
	middlewares []negroni.Handler
}

func (r *Router) Group(path string, fn func(r *Router), middleware ...negroni.Handler) {
	g := group{path: r.sanitize(path)}
	for _, m := range middleware {
		g.middlewares = append(g.middlewares, m)
	}
	r.groups = append(r.groups, g)

	fn(r)
	r.groups = r.groups[:len(r.groups)-1]
}

//////////////////////////////////////////////////////////////////////

func (r *Router) sanitize(path string) string {
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

//////////////////////////////////////////////////////////////////////

const paramsKey int = iota

func Params(r *http.Request) httprouter.Params {
	if rv := httpcontext.Get(r, paramsKey); rv != nil {
		return rv.(httprouter.Params)
	}
	return nil
}

func setParams(r *http.Request, val httprouter.Params) {
	httpcontext.Set(r, paramsKey, val)
}
