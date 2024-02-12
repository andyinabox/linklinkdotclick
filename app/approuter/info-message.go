package approuter

// type InfoPageBody struct {
// 	Message  string
// 	Error    error
// 	LinkUrl  string
// 	LinkText string
// }

// type InfoPageRenderContext struct {
// 	Head HeadRenderContext
// 	Body InfoPageBody
// 	Foot FootRenderContext
// }

// func (r *Router) InfoMessage(ctx *gin.Context, status int, message string, err error, linkUrl string, linkText string) {
// 	ctx.HTML(status, "info.html.tmpl", &InfoPageRenderContext{
// 		r.NewHeadRenderContext(ctx),
// 		InfoPageBody{
// 			Message:  message,
// 			Error:    err,
// 			LinkUrl:  linkUrl,
// 			LinkText: linkText,
// 		},
// 		r.NewFootRenderContext(ctx),
// 	})
// }

// func (r *Router) InfoMessageError(ctx *gin.Context, status int, err error) {
// 	r.InfoMessage(ctx, status, "🫠 Uh-oh, something went wrong...", err, "/", "Back to safety")
// }

// func (r *Router) InfoMessageSuccess(ctx *gin.Context, message string) {
// 	r.InfoMessage(ctx, http.StatusOK, message, nil, "/", "Back to the main page")
// }
