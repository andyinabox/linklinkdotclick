package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexRenderContext struct {
	Links     []Link
	DummyLink Link
}

func (a *App) IndexGet(ctx *gin.Context) {

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
	})
}
