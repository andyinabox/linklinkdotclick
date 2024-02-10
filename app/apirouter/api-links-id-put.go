package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdPut(ctx *gin.Context) {
	logger := r.sc.LogService()

	_, refresh := ctx.GetQuery("refresh")

	id, err := r.hh.GetIdParam(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	var link app.Link

	err = ctx.BindJSON(&link)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	link.UserID = userId

	updatedLink, err := r.sc.LinkService().UpdateLink(userId, id, link, refresh)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.SuccessResponseJSON(ctx, updatedLink)
}
