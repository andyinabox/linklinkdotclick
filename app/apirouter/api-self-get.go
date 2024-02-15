package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiSelfGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hh.SuccessResponseJSON(ctx, user)
}
