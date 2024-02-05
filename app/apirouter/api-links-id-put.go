package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdPut(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

	id, err := r.hh.GetID(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	var link app.Link

	ls, err := r.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ctx.BindJSON(&link)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	updatedLink, err := ls.UpdateLink(id, link, refresh)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.SuccessResponseJSON(ctx, updatedLink)
}
