package apirouter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksPostBody struct {
	Url string
}

func (r *Router) ApiLinksPost(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := ctx.GetUint("userId")

	var body ApiLinksPostBody
	err := ctx.BindJSON(&body)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	if body.Url == "" {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, errors.New("missing url"))
		return
	}

	link, err := r.sc.LinkService().CreateLink(userId, body.Url)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.jrh.ResponseSuccessCreated(ctx, link)

}
