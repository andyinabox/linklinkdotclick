package app

import (
	"embed"
	"fmt"
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
	Host string
	Port string
	Mode string
	// TimeZone  string
	Resources embed.FS
}

func New(conf *Config, ls LinkService) *App {

	// loc, err := time.LoadLocation(conf.TimeZone)
	// if err != nil {
	// 	panic(err)
	// }
	// time.Local = loc

	templates, err := template.ParseFS(conf.Resources, "res/tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	staticFiles, err := fs.Sub(fs.FS(conf.Resources), "res/static")
	if err != nil {
		panic(err)
	}

	gin.SetMode(conf.Mode)
	router := gin.Default()
	router.SetHTMLTemplate(templates)
	router.StaticFS("/static", http.FS(staticFiles))

	app := &App{conf, router, ls}

	router.GET("/", app.IndexGet)

	api := router.Group("/api")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"localhost", "127.0.0.1", "linklink.click"}
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
		return autotls.Run(a.router, "linklink.click")

		// run with custom host/port
	} else if a.conf.Host != "" || a.conf.Port != "" {
		return a.router.Run(fmt.Sprintf("%s:%s", a.conf.Host, a.conf.Port))

		// run with defaults
	} else {
		return a.router.Run()
	}
}
