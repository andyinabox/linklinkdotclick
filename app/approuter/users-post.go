package approuter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) UsersPost(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := r.ah.UserId(ctx)

	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	user.SiteTitle = ctx.PostForm("site-title")

	_, err = r.sc.UserService().UpdateUser(userId, *user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
