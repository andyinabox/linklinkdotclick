package simpleserver

import "net/http"

var defaultMethods = []string{http.MethodGet}

type HandlerFunc func(ctx *Context) http.HandlerFunc

type RouteOptions struct {
	Methods []string
}

// Route creates a new http route
func (s *Server) Route(path string, handler HandlerFunc, opts *RouteOptions) {
	methods := defaultMethods
	if len(opts.Methods) > 0 {
		methods = opts.Methods
	}
	s.router.HandleFunc(path, handler(s.ctx)).
		Methods(methods...)
}
