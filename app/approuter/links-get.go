package approuter

import (
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) LinksGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := ginhelper.GetUint(ctx, "userId")
	id, err := ginhelper.GetQueryUint(ctx, "id")
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
	link.LastClicked = time.Now()
	link, err = ls.UpdateLink(userId, id, *link, false)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, link.SiteUrl)
}
