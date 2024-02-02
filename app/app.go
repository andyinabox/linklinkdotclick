package app

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

type App struct {
	conf   *Config
	router *gin.Engine
	ls     LinkService
}

type Config struct {
	Domain    string
	Port      string
	Mode      string
	Resources embed.FS
}

func New(conf *Config, ls LinkService) *App {

	// load templates
	templates, err := template.ParseFS(conf.Resources, "res/tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	// setup static files filesystem
	staticFiles, err := fs.Sub(fs.FS(conf.Resources), "res/static")
	if err != nil {
		panic(err)
	}

	// configure cors
	corsConfig := cors.DefaultConfig()
	if conf.Mode == "release" {
		corsConfig.AllowOrigins = []string{"https://" + conf.Domain}
	} else {
		corsConfig.AllowOrigins = []string{"http://localhost", "http://127.0.0.1"}
	}

	// router setup
	gin.SetMode(conf.Mode)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.SetHTMLTemplate(templates)

	// create app
	app := &App{conf, router, ls}

	// serve static files
	router.StaticFS("/static", http.FS(staticFiles))

	// http routes
	router.GET("/", app.IndexGet)

	// api routes
	api := router.Group("/api")
	api.Use(cors.New(corsConfig))
	api.GET("/links", app.ApiLinksGet)
	api.POST("/links", app.ApiLinksPost)
	api.GET("/links/:id", app.ApiLinksIdGet)
	api.PUT("/links/:id", app.ApiLinksIdPut)
	api.DELETE("/links/:id", app.ApiLinksIdDelete)

	return app
}

func (a *App) Start() error {

	// run with let's encrypt
	if a.conf.Mode == "release" {
		return autotls.Run(a.router, a.conf.Domain)

		// run with custom host/port
	} else if a.conf.Port != "" {
		return a.router.Run(":" + a.conf.Port)

		// run with defaults
	} else {
		return a.router.Run()
	}
}
