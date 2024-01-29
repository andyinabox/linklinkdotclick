package app

import (
	"embed"

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

	server.Route("/", app.IndexGet, &simpleserver.RouteOptions{})
	server.Route("/api/links", app.ApiLinksGet, &simpleserver.RouteOptions{})
	server.Route("/api/links/{id}", app.ApiLinksIdGet, &simpleserver.RouteOptions{})

	return app
}

func (a *App) Start() error {
	return a.server.Serve()
}
