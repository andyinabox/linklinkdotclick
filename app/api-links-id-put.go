package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ApiLinksIdPut(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

	id, err := a.GetID(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	var link Link

	ls, err := a.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ctx.BindJSON(&link)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	updatedLink, err := ls.UpdateLink(id, link, refresh)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	a.SuccessResponseJSON(ctx, updatedLink)
}
