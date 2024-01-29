package server

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

var defaultMethods = []string{http.MethodGet}

type ServerConfig struct {
	EmbedFS        embed.FS
	Port           string
	Host           string
	TemplatesGlob  string
	StaticDirName  string
	EmbedFSRootDir string
}

type ServerContext struct {
	Resources embed.FS
	Templates *template.Template
}

type Server struct {
	conf   *ServerConfig
	fs     fs.FS
	router *mux.Router
	ctx    *ServerContext
}

type RouteOptions struct {
	Methods []string
}

type HandlerFunc func(ctx *ServerContext) http.HandlerFunc

func NewServer(conf *ServerConfig) *Server {

	// compile templates
	templates, err := template.ParseFS(conf.EmbedFS, conf.TemplatesGlob)
	if err != nil {
		panic(err)
	}

	ctx := &ServerContext{conf.EmbedFS, templates}

	// create fs from embedded files
	fSys, err := fs.Sub(fs.FS(conf.EmbedFS), conf.EmbedFSRootDir)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	return &Server{conf, fSys, router, ctx}
}

func (s *Server) Route(path string, handler HandlerFunc, opts *RouteOptions) {
	methods := defaultMethods
	if len(opts.Methods) > 0 {
		methods = opts.Methods
	}
	s.router.HandleFunc(path, handler(s.ctx)).
		Methods(methods...)
}

func (s *Server) Serve() error {

	// serve embedded static files
	s.router.PathPrefix(s.conf.StaticDirName).Handler(http.FileServer(http.FS(s.fs)))

	// run server
	http.Handle("/", s.router)
	addr := fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port)
	fmt.Printf("Running server on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}
