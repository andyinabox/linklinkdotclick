package approuter

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (r *Router) LoginGet(ctx *gin.Context) {

	hash := ctx.Param("hash")
	user, err := r.sc.UserService().GetUserFromLoginHash(hash)

	if err != nil || user == nil {
		r.InfoMessage(ctx, http.StatusUnauthorized, "🔒 That link isn't working. Did you already use it?", err, "/", "Try sending a new one")
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", user.ID)
	session.Save()

	ctx.Redirect(http.StatusSeeOther, "/")
}
