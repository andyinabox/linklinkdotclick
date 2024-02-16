package approuter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) StylesPost(ctx *gin.Context) {

	userId := ctx.GetUint("userId")
	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
	}

	styles := ctx.PostForm("styles")
	user.StyleSheet = styles
	_, err = r.sc.UserService().UpdateUser(userId, *user)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
