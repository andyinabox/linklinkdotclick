package approuter

import (
	"html/template"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Templates *template.Template
	Version   string
	SmtpAddr  string
}

type Router struct {
	sc   app.ServiceContainer
	hh   app.HandlerHelper
	conf *Config
}

func New(sc app.ServiceContainer, hh app.HandlerHelper, conf *Config) *Router {
	return &Router{sc, hh, conf}
}

func (r *Router) Register(engine *gin.Engine) {
	// main page
	engine.GET("/", r.IndexGet)

	// other pages
	engine.GET("/about", r.AboutGet)

	// auth
	engine.POST("/login", r.LoginPost)
	engine.GET("/login/:hash", r.LoginGet)
	engine.POST("/logout", r.LogoutPost)

	// opml
	engine.GET("/opml", r.OpmlGet)
	engine.POST("/opml", r.OpmlPost)

}
