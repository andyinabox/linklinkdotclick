package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdPut(ctx *gin.Context) {
	logger := r.sc.LogService()

	_, refresh := ctx.GetQuery("refresh")

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	userId, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusUnauthorized, err)
		return
	}

	var link app.Link

	err = ctx.BindJSON(&link)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	link.UserID = userId

	updatedLink, err := r.sc.LinkService().UpdateLink(userId, id, link, refresh)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.jrh.ResponseSuccessPayload(ctx, updatedLink)
}
