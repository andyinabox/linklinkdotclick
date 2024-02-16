package approuter

import (
	"errors"
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

	output, valid, _ := cssparser.Parse([]byte(styles), &cssparser.ParseOptions{})
	if !valid {
		err = errors.New("invalid css")
		r.hrh.PageInfoError(ctx, http.StatusBadRequest, err)
		return
	}

	user.StyleSheet = string(output)
	_, err = r.sc.UserService().UpdateUser(userId, *user)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
