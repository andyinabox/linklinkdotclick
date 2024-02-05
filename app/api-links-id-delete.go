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

	ls, err := a.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	id, err = ls.DeleteLink(id)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	a.SuccessResponseJSON(ctx, &ApiLinksIdDeleteResponse{id})
}
