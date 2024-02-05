package app

import "github.com/gin-gonic/gin"

type HandlerHelper interface {
	// response helpers
	ResponseJSON(ctx *gin.Context, code int, success bool, err error, data interface{})
	ErrorResponse(ctx *gin.Context, code int, err error)
	SuccessResponse(ctx *gin.Context)
	SuccessResponseJSON(ctx *gin.Context, data interface{})
	CreatedResponseJSON(ctx *gin.Context, data interface{})
	NotFoundResponse(ctx *gin.Context)

	// data retrieval
	GetID(ctx *gin.Context) (uint, error)
	GetUserIdFromSession(ctx *gin.Context) (id uint, err error)
	GetUserLinkServiceFromSession(ctx *gin.Context) (LinkService, error)
	GetUserFromSession(ctx *gin.Context) (*User, error)
}
