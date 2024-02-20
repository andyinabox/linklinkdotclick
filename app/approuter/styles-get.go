package approuter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) StylesGet(ctx *gin.Context) {
	user, err := r.ah.User(ctx)
	if err != nil {
		r.hrh.PageInfoError(ctx, http.StatusUnauthorized, err)
		return
	}
	r.hrh.PageStyleEditor(ctx, user)
}
