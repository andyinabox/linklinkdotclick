package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksPostBody struct {
	Url string
}

func (a *App) ApiLinksPost(ctx *gin.Context) {

	var body ApiLinksPostBody
	err := ctx.BindJSON(&body)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if body.Url == "" {
		a.hh.ErrorResponse(ctx, http.StatusBadRequest, errors.New("missing url"))
		return
	}

	ls, err := a.hh.GetUserLinkServiceFromSession(ctx)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	link, err := ls.CreateLink(body.Url)
	if err != nil {
		a.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	a.hh.CreatedResponseJSON(ctx, link)

}
