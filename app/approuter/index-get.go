package approuter

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type IndexRenderContext struct {
	Links     []app.Link
	DummyLink app.Link
	User      *app.User
}

func (r *Router) IndexGet(ctx *gin.Context) {

	user, err := r.hh.GetUserFromSession(ctx)
	// it's ok if no user is found, but we want to abort for server errors
	if err != nil && errors.Is(err, app.ErrServerError) {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ls, err := r.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	links, err := ls.FetchLinks()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "index.html.tmpl", &IndexRenderContext{
		Links: links,
		User:  user,
	})
}