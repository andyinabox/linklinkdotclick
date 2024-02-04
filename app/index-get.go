package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexRenderContext struct {
	Links     []Link
	DummyLink Link
	User      *User
}

func (a *App) IndexGet(ctx *gin.Context) {

	user, err := a.GetUserFromSession(ctx)
	// it's ok if no user is found, but we want to abort for server errors
	if err != nil && errors.Is(err, ErrServerError) {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ls, err := a.GetUserLinkServiceFromSession(ctx)
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
