package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) LinksUpdateIdPost(ctx *gin.Context) {

	logger := r.sc.LogService()

	userId, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusUnauthorized, err)
		return
	}

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	var updates app.Link
	err = ctx.ShouldBind(&updates)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	ls := r.sc.LinkService()

	link, err := ls.FetchLink(userId, id, false)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}
	link.SiteName = updates.SiteName
	link.SiteUrl = updates.SiteUrl
	link.FeedUrl = updates.FeedUrl
	link.HideUnreadCount = updates.HideUnreadCount

	_, err = ls.UpdateLink(userId, id, *link, true)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
