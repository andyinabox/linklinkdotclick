package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksGet(ctx *gin.Context) {

	ls, err := r.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	links, err := ls.FetchLinks()
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hh.SuccessResponseJSON(ctx, links)
}
