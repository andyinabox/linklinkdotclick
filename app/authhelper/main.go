// Package authhelper provides a Gin middleware function to access the User ID from the current session.
// It also provides several helper methods for retrieving auth data within response handlers.
package authhelper

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	userIdKey string = "userId"
)

type Config struct {
	SessionUserKey string
}

type Helper struct {
	us   app.UserService
	conf *Config
}

func New(us app.UserService, conf *Config) *Helper {
	return &Helper{us, conf}
}

func (h *Helper) UserId(ctx *gin.Context) (id uint) {
	return ctx.GetUint(userIdKey)
}

func (h *Helper) User(ctx *gin.Context) (user *app.User, err error) {
	return h.us.FetchUser(h.UserId(ctx))
}

func (h *Helper) IsAuthenticated(ctx *gin.Context) bool {
	id := ctx.GetUint(userIdKey)
	return id != 0
}

func (h *Helper) AuthnMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		value := session.Get(h.conf.SessionUserKey)
		id, ok := value.(uint)
		if ok && id != 0 {
			ctx.Set(userIdKey, id)
		}

		ctx.Next()
	}
}
