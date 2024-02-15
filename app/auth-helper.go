package app

import "github.com/gin-gonic/gin"

type AuthHelper interface {
	AuthMiddleware() gin.HandlerFunc
	GetUserIdFromSession(ctx *gin.Context) (id uint, fakeUser bool, err error)
	GetUserFromSession(ctx *gin.Context) (user *User, fakeUser bool, err error)
}
