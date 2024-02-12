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
	ctx.HTML(status, "info.html.tmpl", h.newRenderContext(ctx, infoPageBody{
		HtmlInfoMessageOptions: *opts,
	}, &headOptions{
		RedirectTimeoutSeconds: opts.RedirectTimeoutSeconds,
		RedirectUrl:            opts.RedirectUrl,
	}))
}

func (h *Helper) InfoPageError(ctx *gin.Context, status int, err error, redirect bool) {
	opts := h.conf.InfoPageErrorOptions
	opts.Error = err
	if redirect {
		opts.RedirectTimeoutSeconds = 3
	}
	h.InfoPage(ctx, status, opts)
}

func (h *Helper) InfoPageSuccess(ctx *gin.Context, message string, redirect bool) {
	opts := h.conf.InfoPageSuccessOptions
	opts.Message = message
	if redirect {
		opts.RedirectTimeoutSeconds = 3
	}
	h.InfoPage(ctx, http.StatusOK, opts)
}
