package approuter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) StylesGet(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusUnauthorized, err)
		return
	}
	r.hrh.PageStyleEditor(ctx, user)
}
