package approuter

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/pushit"
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

	app := engine.Group("")
	app.Use(r.ah.AuthnMiddleware())
	app.Use(pushit.Middleware([]pushit.Resource{
		{
			Path:        "/static/normalize.css",
			ContentType: pushit.ContentTypeStyle,
		},
		{
			Path:        fmt.Sprintf("/static/main.%s.css", r.conf.Version),
			ContentType: pushit.ContentTypeStyle,
		},
		{
			Path:        fmt.Sprintf("/static/main.%s.js", r.conf.Version),
			ContentType: pushit.ContentTypeScript,
		},
	}))

	app.GET("/", r.IndexGet)
	app.GET("/about", r.AboutGet)

	// auth
	app.POST("/session", r.SessionPost)
	app.GET("/session", r.SessionGet)

	requireAuthn := app.Group("", func(ctx *gin.Context) {
		if !r.ah.IsAuthenticated(ctx) {
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
		}
		ctx.Next()
	})

	{
		// styles
		requireAuthn.POST("/styles", r.StylesPost)

		// opml
		requireAuthn.GET("/opml", r.OpmlGet)
		requireAuthn.POST("/opml", r.OpmlPost)

		// links\
		requireAuthn.GET("/links", r.LinksGet)
		requireAuthn.POST("/links", r.LinksPost)

		// users
		requireAuthn.POST("/users", r.UsersPost)
	}

}
