package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiLinksIdClickPut(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := ginhelper.GetUint(ctx, "userId")

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, err)
	}

	link, err := r.sc.LinkService().RegisterLinkClick(userId, id)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
	}

	r.jrh.ResponseSuccessPayload(ctx, link)
}
