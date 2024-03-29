package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

type ApiLinksIdDeleteResponse struct {
	ID uint `json:"id"`
}

func (r *Router) ApiLinksIdDelete(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := r.ah.UserId(ctx)

	id, err := ginhelper.GetParamUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	id, err = r.sc.LinkService().DeleteLink(userId, id)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.jrh.ResponseSuccessPayload(ctx, &ApiLinksIdDeleteResponse{id})
}
