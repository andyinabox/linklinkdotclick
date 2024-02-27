package app

import "github.com/gin-gonic/gin"

type AuthHelper interface {
	AuthnMiddleware() gin.HandlerFunc

	UserId(ctx *gin.Context) (id uint)
	User(ctx *gin.Context) (user *User, err error)
	IsAuthenticated(ctx *gin.Context) bool
}
