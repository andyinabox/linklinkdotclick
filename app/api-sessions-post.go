package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiSessionsPostBody struct {
	Email string
}

func (a *App) ApiSessionsPost(ctx *gin.Context) {
	var body ApiSessionsPostBody
	err := ctx.BindJSON(&body)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if body.Email == "" {
		a.ErrorResponse(ctx, http.StatusBadRequest, errors.New("missing email"))
		return
	}

	token, err := a.us.Login(body.Email)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie("session", token, 3600, "/", ctx.Request.Host, true, true)

	a.SuccessResponse(ctx)

}
