package apirouter

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Domain string
	Mode   string
}

type Router struct {
	sc   app.ServiceContainer
	ah   app.AuthHelper
	jrh  app.JsonResponseHelper
	conf *Config
}

func New(sc app.ServiceContainer, ah app.AuthHelper, jrh app.JsonResponseHelper, conf *Config) *Router {
	router := &Router{sc, ah, jrh, conf}

	return router
}

func (r *Router) Register(engine *gin.Engine) {

	corsConfig := cors.DefaultConfig()
	if r.conf.Mode == "release" {
		corsConfig.AllowOrigins = []string{"https://" + r.conf.Domain}
	} else {
		corsConfig.AllowOrigins = []string{"http://localhost", "http://127.0.0.1"}
	}

	api := engine.Group("/api")
	api.Use(cors.New(corsConfig))
	api.Use(r.ah.AuthnMiddleware())
	// error if not authenticated
	api.Use(func(ctx *gin.Context) {
		if !r.ah.IsAuthenticated(ctx) {
			r.jrh.ResponseError(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
			ctx.Abort()
		}
	})

	// links
	api.GET("/links", r.ApiLinksGet)
	api.POST("/links", r.ApiLinksPost)
	api.GET("/links/:id", r.ApiLinksIdGet)
	api.PUT("/links/:id", r.ApiLinksIdPut)
	api.DELETE("/links/:id", r.ApiLinksIdDelete)
	api.PATCH("/links/:id", r.ApiLinksIdPatch)

}
