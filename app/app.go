package app

import (
	"embed"

	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

type App struct {
	conf   *Config
	server *simpleserver.Server
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

	app := &App{conf, server}

	server.Route("/", app.GetIndex, &simpleserver.RouteOptions{})

	return app
}

func (a *App) Start() error {
	return a.server.Serve()
}
