package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksIdDeleteResponse struct {
	ID uint `json:"id"`
}

func (a *App) ApiLinksIdDelete(ctx *gin.Context) {
	id, err := a.GetID(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	tx := a.db.Delete(&Link{}, id)
	err = tx.Error
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	a.SuccessResponseJSON(ctx, &ApiLinksIdDeleteResponse{id})
}
