package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/cssparser"
	"github.com/gin-gonic/gin"
)

func (r *Router) StylesPost(ctx *gin.Context) {

	userId := ctx.GetUint("userId")
	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	styles := ctx.PostForm("styles")

	result, err := cssparser.Parse([]byte(styles), true)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	user.StyleSheet = string(result.Output)
	_, err = r.sc.UserService().UpdateUser(userId, *user)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
