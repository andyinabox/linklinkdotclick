package approuter

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (r *Router) LogoutPost(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/")
}
