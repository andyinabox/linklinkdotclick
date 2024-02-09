package approuter

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (r *Router) LogoutPost(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	r.InfoMessageSuccess(ctx, "👋 You're logged out. Later!")
}
