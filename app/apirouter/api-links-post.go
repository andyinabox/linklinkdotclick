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

	ls, err := r.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	link, err := ls.CreateLink(body.Url)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.CreatedResponseJSON(ctx, link)

}
