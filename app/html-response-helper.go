package app

import "github.com/gin-gonic/gin"

type HtmlInfoMessageOptions struct {
	Message  string
	Error    error
	LinkText string
	LinkUrl  string
}

type HtmlResponseHelper interface {
	HomePage(ctx *gin.Context, user *User, isDefaultUser bool, links []Link, editMode bool)
	InfoPage(ctx *gin.Context, status int, opts *HtmlInfoMessageOptions)
	InfoPageError(ctx *gin.Context, status int, err error)
	InfoPageSuccess(ctx *gin.Context, message string)
	AboutPage(ctx *gin.Context)
}
