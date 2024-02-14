package authhelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Config struct {
	SessionUserKey string
}

type Helper struct {
	sc   app.ServiceContainer
	conf *Config
}

func New(sc app.ServiceContainer, conf *Config) *Helper {
	return &Helper{sc, conf}
}

func (h *Helper) GetUserIdFromSession(ctx *gin.Context) (id uint, isDefaultUser bool, err error) {
	session := sessions.Default(ctx)
	value := session.Get(h.conf.SessionUserKey)
	if value == nil {
		return h.sc.UserService().GetDefaultUserId(), true, nil
	}
	var ok bool
	id, ok = value.(uint)
	if !ok {
		err = app.ErrServerError
		return
	}
	return
}

func (h *Helper) GetUserFromSession(ctx *gin.Context) (*app.User, bool, error) {
	id, isDefaultUser, err := h.GetUserIdFromSession(ctx)
	if err != nil {
		return nil, false, err
	}
	user, err := h.sc.UserService().FetchUser(id)
	return user, isDefaultUser, err
}

func (h *Helper) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId, isDefaultUser, err := h.GetUserIdFromSession(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}

		ctx.Set("userId", userId)
		ctx.Set("isDefaultUser", isDefaultUser)

		ctx.Next()

		// I don't think there's anything I need to do here
	}
}
