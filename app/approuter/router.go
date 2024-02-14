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
	ah   app.AuthHelper
	hrh  app.HtmlResponseHelper
	conf *Config
}

func New(sc app.ServiceContainer, ah app.AuthHelper, hrh app.HtmlResponseHelper, conf *Config) *Router {
	return &Router{sc, ah, hrh, conf}
}

func (r *Router) Register(engine *gin.Engine) {

	app := engine.Group("/")
	app.Use(r.ah.AuthMiddleware())

	// main page
	app.GET("/", r.IndexGet)

	// other pages
	app.GET("/about", r.AboutGet)

	// auth
	app.POST("/session", r.SessionPost)
	app.GET("/session", r.SessionGet)

	// opml
	app.GET("/opml", r.OpmlGet)
	app.POST("/opml", r.OpmlPost)

	// links\
	app.GET("/links", r.LinksGet)
	app.POST("/links", r.LinksPost)

	// users
	app.POST("/users", r.UsersPost)

}
