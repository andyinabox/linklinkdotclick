package apirouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ApiSelfGet(ctx *gin.Context) {

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	user, err := r.sc.UserService().FetchUser(userId)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hh.SuccessResponseJSON(ctx, user)
}
