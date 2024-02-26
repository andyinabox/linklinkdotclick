package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) IndexGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	if !r.ah.IsAuthenticated(ctx) {
		r.hrh.PageAbout(ctx)
		return
	}

	isEditing := ginhelper.GetQueryBool(ctx, "editing")
	user, err := r.ah.User(ctx)

	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(user.ID)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hrh.PageHome(ctx, user, false, links, isEditing)
}
