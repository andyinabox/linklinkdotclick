package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdPut(ctx *gin.Context) {
	logger := r.sc.LogService()

	_, refresh := ctx.GetQuery("refresh")

	userId := ctx.GetUint("userId")

	var link app.Link

	err := ctx.BindJSON(&link)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	link.UserID = userId

	updatedLink, err := r.sc.LinkService().UpdateLink(userId, link, refresh)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.jrh.ResponseSuccessPayload(ctx, updatedLink)
}
