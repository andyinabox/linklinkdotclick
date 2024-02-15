package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type infoPageBody struct {
	app.HtmlInfoMessageOptions
}

func (h *Helper) InfoPage(ctx *gin.Context, status int, opts *app.HtmlInfoMessageOptions) {
	ctx.HTML(status, "base.html.tmpl", h.newRenderContext(ctx, info, infoPageBody{
		HtmlInfoMessageOptions: *opts,
	}))
}

func (h *Helper) InfoPageError(ctx *gin.Context, status int, err error) {
	opts := *h.conf.InfoPageErrorOptions
	opts.Error = err
	h.InfoPage(ctx, status, &opts)
}

func (h *Helper) InfoPageSuccess(ctx *gin.Context, message string) {
	opts := *h.conf.InfoPageSuccessOptions
	opts.Message = message
	h.InfoPage(ctx, http.StatusOK, &opts)
}
