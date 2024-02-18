package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/cssparser"
	"github.com/gin-gonic/gin"
)

func (r *Router) StylesPost(ctx *gin.Context) {

	user, err := r.ah.User(ctx)
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
	_, err = r.sc.UserService().UpdateUser(user.ID, *user)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
