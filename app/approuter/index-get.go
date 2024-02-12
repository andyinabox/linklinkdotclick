package approuter

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) IndexGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	isEditing := ginhelper.GetQueryBool(ctx, "editing")

	user, isDefaultUser, err := r.ah.GetUserFromSession(ctx)
	if err != nil && errors.Is(err, app.ErrServerError) {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err, false)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(user.ID)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err, false)
		return
	}

	r.hrh.HomePage(ctx, user, isDefaultUser, links, isEditing)
}
