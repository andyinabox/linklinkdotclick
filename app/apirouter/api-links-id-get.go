package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdGet(ctx *gin.Context) {
	logger := r.sc.LogService()
	_, refresh := ctx.GetQuery("refresh")

	userId := ctx.GetUint("userId")

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	link, err := r.sc.LinkService().FetchLink(userId, id, refresh)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.jrh.ResponseSuccessPayload(ctx, link)
}
