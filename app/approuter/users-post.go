package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

func (r *Router) UsersPost(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := r.ah.UserId(ctx)

	var user app.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}
	user.ID = userId

	_, err = r.sc.UserService().UpdateUser(userId, user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
