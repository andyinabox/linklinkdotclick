package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ApiLinksGet(ctx *gin.Context) {

	ls, err := a.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	links, err := ls.FetchLinks()
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	a.hh.SuccessResponseJSON(ctx, links)
}
