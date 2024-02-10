package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiSelfPut(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	var user app.User

	err = ctx.BindJSON(&user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	updatedUser, err := r.sc.UserService().UpdateUser(userId, user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.SuccessResponseJSON(ctx, updatedUser)
}
