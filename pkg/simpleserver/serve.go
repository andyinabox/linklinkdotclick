package simpleserver

import (
	"fmt"
	"net/http"
)

// Serve starts the simple server
func (s *Server) Serve() error {

	// serve embedded static files
	s.router.PathPrefix(s.conf.StaticDirName).Handler(http.FileServer(http.FS(s.fs)))

	// run server
	http.Handle("/", s.router)
	addr := fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port)
	fmt.Printf("Running server on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}
