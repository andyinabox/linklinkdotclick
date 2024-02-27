package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type pageInfoBody struct {
	app.HtmlInfoMessageOptions
}

func (h *Helper) PageInfo(ctx *gin.Context, status int, opts *app.HtmlInfoMessageOptions) {
	ctx.HTML(status, "base.html.tmpl", h.newRenderContext(ctx, pageInfo, pageInfoBody{
		HtmlInfoMessageOptions: *opts,
	}))
}

func (h *Helper) PageInfoError(ctx *gin.Context, status int, err error) {
	opts := *h.conf.InfoPageErrorOptions
	opts.Error = err
	h.PageInfo(ctx, status, &opts)
}

func (h *Helper) PageInfoSuccess(ctx *gin.Context, message string) {
	opts := *h.conf.InfoPageSuccessOptions
	opts.Message = message
	h.PageInfo(ctx, http.StatusOK, &opts)
}
