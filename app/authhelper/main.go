// Package authhelper provides a Gin middleware function to access the User ID from the current session.
// It also provides several helper methods for retrieving auth data within response handlers.
package authhelper

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	userIdKey        string = "userId"
	isDefaultUserKey string = "isDefaultUser"
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

func (h *Helper) IsDefaultUser(ctx *gin.Context) bool {
	return ctx.GetBool(isDefaultUserKey)
}

func (h *Helper) User(ctx *gin.Context) (user *app.User, err error) {
	return h.us.FetchUser(h.UserId(ctx))
}

func (h *Helper) getUserIdFromSession(ctx *gin.Context) (id uint, isDefaultUser bool, err error) {
	session := sessions.Default(ctx)
	value := session.Get(h.conf.SessionUserKey)
	if value == nil {
		return h.us.GetDefaultUserId(), true, nil
	}

	var ok bool
	if id, ok = value.(uint); !ok {
		err = errors.New("unable to type user id as uint")
		return
	}

	return
}

func (h *Helper) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId, isDefaultUser, err := h.getUserIdFromSession(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		ctx.Set(userIdKey, userId)
		ctx.Set(isDefaultUserKey, isDefaultUser)

		ctx.Next()

	}
}
