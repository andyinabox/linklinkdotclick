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

	var body ApiLinksPostBody
	err := ctx.BindJSON(&body)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if body.Url == "" {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, errors.New("missing url"))
		return
	}

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	link, err := r.sc.LinkService().CreateLink(userId, body.Url)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.CreatedResponseJSON(ctx, link)

}
