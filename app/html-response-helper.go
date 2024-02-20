package app

import "github.com/gin-gonic/gin"

type HtmlInfoMessageOptions struct {
	Message  string
	Error    error
	LinkText string
	LinkUrl  string
}

type HtmlResponseHelper interface {
	PageHome(ctx *gin.Context, user *User, isDefaultUser bool, links []Link, editMode bool)
	PageAbout(ctx *gin.Context)
	PageInfo(ctx *gin.Context, status int, opts *HtmlInfoMessageOptions)
	PageInfoError(ctx *gin.Context, status int, err error)
	PageInfoSuccess(ctx *gin.Context, message string)
	PageStyleEditor(ctx *gin.Context, user *User)
}
