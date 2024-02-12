package approuter

import (
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) LinksIdGet(ctx *gin.Context) {
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
