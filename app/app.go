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
<<<<<<< HEAD
	engine *gin.Engine
=======
	router *gin.Engine
	us     UserService
	ls     LinkService
>>>>>>> main
}

type Config struct {
	Domain    string
	Port      string
	Mode      string
	Resources embed.FS
}

<<<<<<< HEAD
func New(
	store sessions.Store,
	routers []RouterGroup,
	conf *Config,
) *App {
=======
func New(conf *Config, us UserService, store sessions.Store) *App {

	user, err := us.EnsureDefaultUser()
	if err != nil {
		panic(err)
	}
	ls, err := us.GetUserLinkService(user)
	if err != nil {
		panic(err)
	}
>>>>>>> main

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
<<<<<<< HEAD
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.SetHTMLTemplate(templates)
	engine.Use(sessions.Sessions("session", store))

	// serve static files
	engine.StaticFS("/static", http.FS(staticFiles))
=======
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.SetHTMLTemplate(templates)
	router.Use(sessions.Sessions("session", store))

	// create app
	app := &App{conf, router, us, ls}

	// serve static files
	router.StaticFS("/static", http.FS(staticFiles))

	// http routes
	router.GET("/", app.IndexGet)
	router.POST("/login", app.LoginPost)
	router.POST("/logout", app.LogoutPost)
>>>>>>> main

	for _, group := range routers {
		group.Register(engine)
	}

	return &App{conf, engine}
}

func (a *App) Start() error {

	// run with let's encrypt
	if a.conf.Mode == "release" {
<<<<<<< HEAD
		return autotls.Run(a.engine, a.conf.Domain)
	}

	return a.engine.RunTLS(":"+a.conf.Port, "./.cert/localhost.crt", "./.cert/localhost.key")
=======
		return autotls.Run(a.router, a.conf.Domain)
	}

	return a.router.RunTLS(":"+a.conf.Port, "./.cert/localhost.crt", "./.cert/localhost.key")
>>>>>>> main
}
