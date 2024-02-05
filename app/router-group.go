package app

import "github.com/gin-gonic/gin"

type RouterGroup interface {
	Register(*gin.Engine)
}
