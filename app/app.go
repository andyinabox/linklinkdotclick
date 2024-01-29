package app

import (
	"embed"
	"net/http"

	"github.com/andyinabox/linkydink-sketch/pkg/feedreader"
	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

type App struct {
	conf       *Config
	server     *simpleserver.Server
	feedreader *feedreader.Reader
}

type Config struct {
	Host      string
	Port      string
	Resources embed.FS
}

func New(conf *Config) *App {
	server := simpleserver.NewServer(&simpleserver.Config{
		Host:           conf.Host,
		Port:           conf.Port,
		Resources:      conf.Resources,
		TemplatesGlob:  "res/tmpl/*.tmpl",
		StaticDirName:  "/static/",
		EmbedFSRootDir: "res",
	})

	app := &App{conf, server, &feedreader.Reader{}}

	// main page
	server.Route("/", app.IndexGet, &simpleserver.RouteOptions{})

	// api
	server.Route("/api/links", app.ApiLinksGet, &simpleserver.RouteOptions{
		Methods: []string{http.MethodGet},
	})
	server.Route("/api/links", app.ApiLinksPost, &simpleserver.RouteOptions{
		Methods: []string{http.MethodPost},
	})
	server.Route("/api/links/{id}", app.ApiLinksIdGet, &simpleserver.RouteOptions{
		Methods: []string{http.MethodGet},
	})
	server.Route("/api/links/{id}", app.ApiLinksIdPut, &simpleserver.RouteOptions{
		Methods: []string{http.MethodPut},
	})
	server.Route("/api/links/{id}", app.ApiLinksIdPatch, &simpleserver.RouteOptions{
		Methods: []string{http.MethodPatch},
	})
	server.Route("/api/links/{id}", app.ApiLinksIdDelete, &simpleserver.RouteOptions{
		Methods: []string{http.MethodDelete},
	})

	return app
}

func (a *App) Start() error {
	return a.server.Serve()
}
