package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksIdDeleteResponse struct {
	ID uint `json:"id"`
}

func (a *App) ApiLinksIdDelete(ctx *gin.Context) {
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

	id, err = ls.DeleteLink(id)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	a.hh.SuccessResponseJSON(ctx, &ApiLinksIdDeleteResponse{id})
}
