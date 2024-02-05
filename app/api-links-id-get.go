package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ApiLinksIdGet(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

	id, err := a.hh.GetID(ctx)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ls, err := a.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	link, err := ls.FetchLink(id, refresh)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	a.hh.SuccessResponseJSON(ctx, link)
}
