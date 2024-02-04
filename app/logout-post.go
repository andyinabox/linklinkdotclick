package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) LogoutPost(ctx *gin.Context) {

	token, err := ctx.Cookie("session")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	err = a.us.Logout(token)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	}

	ctx.SetCookie("session", "", -1, "/", "localhost", true, true)
	ctx.Redirect(http.StatusSeeOther, "/")
}
