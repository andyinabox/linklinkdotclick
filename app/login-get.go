package app

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (a *App) LoginGet(ctx *gin.Context) {
	hash := ctx.Param("hash")
	user, err := a.sc.UserService().GetUserFromLoginHash(hash)

	if err != nil || user == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", user.ID)
	session.Save()

	ctx.Redirect(http.StatusSeeOther, "/")
}
