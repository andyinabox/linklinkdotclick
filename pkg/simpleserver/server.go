package simpleserver

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/gorilla/mux"
)

type Config struct {
	Resources      embed.FS
	Port           string
	Host           string
	TemplatesGlob  string
	StaticDirName  string
	EmbedFSRootDir string
}

type Server struct {
	conf   *Config
	fs     fs.FS
	router *mux.Router
	ctx    *Context
}

func New(conf *Config) *Server {

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
		WriteError: WriteError,
		WriteJSON:  WriteJSON,
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
