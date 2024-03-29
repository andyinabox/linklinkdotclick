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

const staticFilesPath = "res/static"

type App struct {
	conf   *Config
	engine *gin.Engine
}

type Config struct {
	Domain      string
	Port        string
	Mode        string
	Resources   embed.FS
	Templates   *template.Template
	Version     string
	SessionName string
}

func New(
	store sessions.Store,
	routers []RouterGroup,
	conf *Config,
) *App {

	// setup static files filesystem
	staticFiles, err := fs.Sub(fs.FS(conf.Resources), staticFilesPath)
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
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.SetHTMLTemplate(conf.Templates)
	engine.Use(sessions.Sessions(conf.SessionName, store))

	// serve static files
	engine.StaticFS("/static", http.FS(staticFiles))

	for _, group := range routers {
		group.Register(engine)
	}

	return &App{conf, engine}
}

func (a *App) Start() error {

	// run with let's encrypt
	if a.conf.Mode == "release" {
		return autotls.Run(a.engine, a.conf.Domain)
	}

	return a.engine.RunTLS(":"+a.conf.Port, "./.cert/localhost.crt", "./.cert/localhost.key")
}
