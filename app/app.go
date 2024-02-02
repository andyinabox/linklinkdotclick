package app

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	conf   *Config
	router *gin.Engine
	ls     LinkService
}

type Config struct {
	// Host      string
	// Port      string
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
	api.GET("/links", app.ApiLinksGet)
	api.POST("/links", app.ApiLinksPost)
	api.GET("/links/:id", app.ApiLinksIdGet)
	api.PUT("/links/:id", app.ApiLinksIdPut)
	api.DELETE("/links/:id", app.ApiLinksIdDelete)

	return app
}

func (a *App) Start() error {
	// return a.router.Run(fmt.Sprintf("%s:%s", a.conf.Host, a.conf.Port))
	return a.router.Run()
}
