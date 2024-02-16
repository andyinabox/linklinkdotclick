package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) IndexGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	isEditing := ginhelper.GetQueryBool(ctx, "editing")
	userId := ctx.GetUint("userId")
	isDefaultUser := ctx.GetBool("isDefaultUser")

	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(userId)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hrh.PageHome(ctx, user, isDefaultUser, links, isEditing)
}
