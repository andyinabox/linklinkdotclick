package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) LinksDeleteIdPost(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusUnauthorized, err, false)
		return
	}

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err, false)
		return
	}

	_, err = r.sc.LinkService().DeleteLink(userId, id)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err, false)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
