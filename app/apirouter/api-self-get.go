package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiSelfGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusUnauthorized, err)
		return
	}

	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		logger.Error().Println(err.Error())
		r.jrh.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.jrh.ResponseSuccessPayload(ctx, user)
}
