package app

import "github.com/gin-gonic/gin"

type JsonResponseHelper interface {
	Response(ctx *gin.Context, code int, success bool, err error, data interface{})
	ResponseError(ctx *gin.Context, code int, err error)
	ResponseSuccess(ctx *gin.Context)
	ResponseSuccessPayload(ctx *gin.Context, data interface{})
	ResponseSuccessCreated(ctx *gin.Context, data interface{})
	ResponseNotFound(ctx *gin.Context)
}
