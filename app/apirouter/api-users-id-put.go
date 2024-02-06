package apirouter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

func (r *Router) ApiUsersIdPut(ctx *gin.Context) {
	id, err := r.hh.GetIdParam(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	userId, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	// user should only be able to modify their own data
	if id != userId {
		r.hh.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	var user app.User

	err = ctx.BindJSON(&user)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	updatedUser, err := r.sc.UserService().UpdateUser(id, user)
	if err != nil {
		r.hh.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	r.hh.SuccessResponseJSON(ctx, updatedUser)
}
