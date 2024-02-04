package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ApiLinksIdGet(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

	id, err := a.GetID(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ls, err := a.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	link, err := ls.FetchLink(id, refresh)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	a.SuccessResponseJSON(ctx, link)
}
