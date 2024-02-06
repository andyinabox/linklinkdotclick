package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksIdDeleteResponse struct {
	ID uint `json:"id"`
}

func (r *Router) ApiLinksIdDelete(ctx *gin.Context) {
	id, err := r.hh.GetIdParam(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	userId, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	id, err = r.sc.LinkService().DeleteLink(userId, id)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hh.SuccessResponseJSON(ctx, &ApiLinksIdDeleteResponse{id})
}
