package apirouter

import (
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

type linkPatch struct {
	LastClicked time.Time `json:"lastClicked"`
}

func (r *Router) ApiLinksIdPatch(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := r.ah.UserId(ctx)

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	var patch linkPatch
	err = ctx.BindJSON(&patch)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	link, err := r.sc.LinkService().RegisterLinkClick(userId, id, patch.LastClicked)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.jrh.ResponseSuccessPayload(ctx, link)
}
