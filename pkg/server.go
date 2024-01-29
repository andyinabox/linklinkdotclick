package pkg

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Res           embed.FS
	Port          string
	Host          string
	TemplatesGlob string
}

type Server struct {
	conf      *ServerConfig
	fs        fs.FS
	templates *template.Template
}

func NewServer(conf *ServerConfig) *Server {

	// compile templates
	templates, err := template.ParseFS(conf.Res, conf.TemplatesGlob)
	if err != nil {
		panic(err)
	}

	// create fs from embedded files
	fSys, err := fs.Sub(fs.FS(conf.Res), "res")
	if err != nil {
		panic(err)
	}

	return &Server{conf, fSys, templates}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	// server index
	router.HandleFunc("/", s.GetIndex).Methods(http.MethodGet)

	// serve static files
	router.PathPrefix("/static/").Handler(http.FileServer(http.FS(s.fs)))

	// run server
	http.Handle("/", router)
	addr := fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port)
	fmt.Printf("Running server on %s\n", addr)
	return http.ListenAndServe(addr, router)
}
