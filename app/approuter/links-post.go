package approuter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) LinksPost(ctx *gin.Context) {
	logger := r.sc.LogService()
	var err error
	originalUrl := ctx.PostForm("url")

	if originalUrl == "" {
		err = errors.New("no url provided")
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	userId, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusUnauthorized, err)
		return
	}

	link, err := r.sc.LinkService().CreateLink(userId, originalUrl)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hrh.InfoPageSuccess(ctx, "✅ Successfully added link "+link.SiteName, false)
}
