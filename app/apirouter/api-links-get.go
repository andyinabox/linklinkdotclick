package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	id, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusUnauthorized, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(id)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.jrh.ResponseSuccessPayload(ctx, links)
}
