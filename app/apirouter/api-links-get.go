package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	id, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(id)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hh.SuccessResponseJSON(ctx, links)
}
