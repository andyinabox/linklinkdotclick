package approuter

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (r *Router) LoginGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	hash := ctx.Param("hash")
	user, err := r.sc.UserService().GetUserFromLoginHash(hash)

	if err == nil && user == nil {
		err = errors.New("no user found for hash")
	}

	if err != nil {
		logger.Error().Println(err.Error())
		r.InfoMessage(ctx, http.StatusUnauthorized, "🔒 That link isn't working. Did you already use it?", err, "/", "Try sending a new one")
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", user.ID)
	session.Save()

	ctx.Redirect(http.StatusSeeOther, "/")
}
