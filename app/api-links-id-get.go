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

	var link Link
	tx := a.db.First(&link, id)
	err = tx.Error
	if err != nil {
		a.NotFoundResponse(ctx)
		return
	}

	if refresh {
		err = a.RefreshLink(&link)
		if err != nil {
			a.ErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	// send response
	a.SuccessResponseJSON(ctx, link)
}
