package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := r.ah.UserId(ctx)

	links, err := r.sc.LinkService().FetchLinks(userId)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.jrh.ResponseSuccessPayload(ctx, links)
}
