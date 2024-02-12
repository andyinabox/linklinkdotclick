package approuter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) SelfUpdatePost(ctx *gin.Context) {
	logger := r.sc.LogService()

	user, _, err := r.ah.GetUserFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusUnauthorized, err)
		return
	}

	siteTitle := ctx.PostForm("site_title")
	if siteTitle != "" {
		user.SiteTitle = siteTitle
	}

	_, err = r.sc.UserService().UpdateUser(user.ID, *user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hrh.InfoPageSuccess(ctx, "✅ Successfully updated your settings", true)

}
