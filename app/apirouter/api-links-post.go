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

	userId, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusUnauthorized, err)
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
