package approuter

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type linksGetQuery struct {
	ID      uint   `form:"id" binding:"required"`
	SiteUrl string `form:"url" binding:"required"`
}

func (r *Router) LinksGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := ctx.GetUint("userId")

	var query linksGetQuery
	err := ctx.BindQuery(&query)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	go r.registerLinkClick(userId, query.ID)

	ctx.Redirect(http.StatusSeeOther, query.SiteUrl)
}

func (r *Router) registerLinkClick(userId uint, id uint) {
	logger := r.sc.LogService()

	_, err := r.sc.LinkService().RegisterLinkClick(userId, id, time.Now())
	if err != nil {
		logger.Error().Println(err.Error())
	}

}
