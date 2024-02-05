package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksIdDeleteResponse struct {
	ID uint `json:"id"`
}

func (r *Router) ApiLinksIdDelete(ctx *gin.Context) {
	id, err := r.hh.GetID(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ls, err := r.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	id, err = ls.DeleteLink(id)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hh.SuccessResponseJSON(ctx, &ApiLinksIdDeleteResponse{id})
}
