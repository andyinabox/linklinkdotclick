package approuter

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type Router struct {
	sc app.ServiceContainer
	hh app.HandlerHelper
}

func New(sc app.ServiceContainer, hh app.HandlerHelper) *Router {
	return &Router{sc, hh}
}

func (r *Router) Register(engine *gin.Engine) {
	engine.GET("/", r.IndexGet)
	engine.POST("/login", r.LoginPost)
	engine.GET("/login/:hash", r.LoginGet)
	engine.POST("/logout", r.LogoutPost)
}
