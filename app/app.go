package app

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

type App struct {
	conf   *Config
	router *gin.Engine
	sc     ServiceContainer
	hh     HandlerHelper
}

type Config struct {
	Domain    string
	Port      string
	Mode      string
	Resources embed.FS
}

func New(conf *Config, sc ServiceContainer, hh HandlerHelper, store sessions.Store) *App {

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
	router.Use(sessions.Sessions("session", store))

	// create app
	app := &App{conf, router, sc, hh}

	// serve static files
	router.StaticFS("/static", http.FS(staticFiles))

	// http routes
	router.GET("/", app.IndexGet)
	router.POST("/login", app.LoginPost)
	router.GET("/login/:hash", app.LoginGet)
	router.POST("/logout", app.LogoutPost)

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
	}

	return a.router.RunTLS(":"+a.conf.Port, "./.cert/localhost.crt", "./.cert/localhost.key")
}
