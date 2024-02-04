package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) LoginPost(ctx *gin.Context) {
	email := ctx.PostForm("email")

	token, err := a.us.Login(email)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	ctx.SetCookie("session", token, 3600, "/", ctx.Request.Host, true, true)
	ctx.SetSameSite(http.SameSiteDefaultMode)
	ctx.Redirect(http.StatusSeeOther, "/")
}
