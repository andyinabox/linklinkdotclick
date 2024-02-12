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
	// main page
	engine.GET("/", r.IndexGet)

	// other pages
	engine.GET("/about", r.AboutGet)

	// auth
	engine.POST("/session", r.SessionPost)
	engine.GET("/session/:hash", r.SessionGetHash)
	engine.POST("/session/delete", r.SessionDelete)

	// opml
	engine.GET("/opml", r.OpmlGet)
	engine.POST("/opml", r.OpmlPost)

	// links
	engine.POST("/links/update/:id", r.LinksUpdateIdPost)
	engine.POST("/links/delete/:id", r.LinksDeleteIdPost)
	engine.GET("/links/:id", r.LinksIdGet)
	engine.POST("/links", r.LinksPost)

	// users
	engine.POST("/self/update", r.SelfUpdatePost)

}
