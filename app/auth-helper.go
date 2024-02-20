package app

import "github.com/gin-gonic/gin"

type AuthHelper interface {
	AuthMiddleware() gin.HandlerFunc
	UserId(ctx *gin.Context) (id uint)
	User(ctx *gin.Context) (user *User, err error)
	IsDefaultUser(ctx *gin.Context) bool
}
