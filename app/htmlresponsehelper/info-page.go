package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type InfoPageBody struct {
	app.HtmlInfoMessageOptions
}

type InfoPageRenderContext struct {
	Head HeadRenderContext
	Body InfoPageBody
	Foot FootRenderContext
}

func (h *Helper) InfoPage(ctx *gin.Context, status int, opts *app.HtmlInfoMessageOptions) {
	ctx.HTML(status, "info.html.tmpl", &InfoPageRenderContext{
		h.NewHeadRenderContext(ctx),
		InfoPageBody{
			HtmlInfoMessageOptions: *opts,
		},
		h.NewFootRenderContext(ctx),
	})
}

func (h *Helper) InfoPageError(ctx *gin.Context, status int, err error) {
	opts := h.conf.InfoPageErrorOptions
	opts.Error = err
	h.InfoPage(ctx, status, opts)
}

func (h *Helper) InfoPageSuccess(ctx *gin.Context, message string) {
	opts := h.conf.InfoPageSuccessOptions
	h.InfoPage(ctx, http.StatusOK, opts)
}
