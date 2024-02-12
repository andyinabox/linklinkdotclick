package app

import "github.com/gin-gonic/gin"

type HtmlInfoMessageOptions struct {
	Message                string
	Error                  error
	RedirectTimeoutSeconds int
	RedirectUrl            string
	LinkText               string
	LinkUrl                string
}

type HtmlResponseHelper interface {
	HomePage(ctx *gin.Context, user *User, isDefaultUser bool, links []Link)
	InfoPage(ctx *gin.Context, status int, opts *HtmlInfoMessageOptions)
	InfoPageError(ctx *gin.Context, status int, err error, redirect bool)
	InfoPageSuccess(ctx *gin.Context, message string, redirect bool)
	AboutPage(ctx *gin.Context)
}
