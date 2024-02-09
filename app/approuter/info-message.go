package approuter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) InfoMessage(ctx *gin.Context, status int, message string, err error, linkUrl string, linkText string) {
	ctx.HTML(status, "info.html.tmpl", &InfoPageRenderContext{
		NewHeadRenderContext(ctx),
		InfoPageBody{
			Message:  message,
			Error:    err,
			LinkUrl:  linkUrl,
			LinkText: linkText,
		},
	})
}

func (r *Router) InfoMessageError(ctx *gin.Context, status int, err error) {
	r.InfoMessage(ctx, status, "🫠 Uh-oh, something went wrong...", err, "/", "Back to safety")
}

func (r *Router) InfoMessageSuccess(ctx *gin.Context, message string) {
	r.InfoMessage(ctx, http.StatusOK, message, nil, "/", "Back to the main page")
}
