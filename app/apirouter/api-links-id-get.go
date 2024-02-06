package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdGet(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

	id, err := r.hh.GetIdParam(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	link, err := r.sc.LinkService().FetchLink(userId, id, refresh)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.SuccessResponseJSON(ctx, link)
}
