package simpleserver

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

var defaultMethods = []string{http.MethodGet}

type Config struct {
	Resources      embed.FS
	Port           string
	Host           string
	TemplatesGlob  string
	StaticDirName  string
	EmbedFSRootDir string
}

type Context struct {
	Resources  embed.FS
	Templates  *template.Template
	Vars       func(*http.Request) map[string]string
	SetUrlVars func(*http.Request, map[string]string) *http.Request
}

type Server struct {
	conf   *Config
	fs     fs.FS
	router *mux.Router
	ctx    *Context
}

type RouteOptions struct {
	Methods []string
}

type HandlerFunc func(ctx *Context) http.HandlerFunc

func NewServer(conf *Config) *Server {

	// compile templates
	templates, err := template.ParseFS(conf.Resources, conf.TemplatesGlob)
	if err != nil {
		panic(err)
	}

	// build context
	ctx := &Context{
		Resources:  conf.Resources,
		Templates:  templates,
		Vars:       mux.Vars,
		SetUrlVars: mux.SetURLVars,
	}

	// create fs from embedded files
	fSys, err := fs.Sub(fs.FS(conf.Resources), conf.EmbedFSRootDir)
	if err != nil {
		panic(err)
	}

	// create router
	router := mux.NewRouter()

	return &Server{conf, fSys, router, ctx}
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
