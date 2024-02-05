package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdGet(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

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

	link, err := ls.FetchLink(id, refresh)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.SuccessResponseJSON(ctx, link)
}
