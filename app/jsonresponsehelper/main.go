package jsonresponsehelper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Helper struct {
}

func New() *Helper {
	return &Helper{}
}

type JSONResponseBody struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func (h *Helper) Response(ctx *gin.Context, code int, success bool, err error, data interface{}) {

	response := &JSONResponseBody{
		Success: success,
		Payload: data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	ctx.JSON(code, response)
}

func (h *Helper) ResponseError(ctx *gin.Context, code int, err error) {
	h.Response(ctx, code, false, err, nil)
}

func (h *Helper) ResponseSuccess(ctx *gin.Context) {
	h.Response(ctx, http.StatusOK, true, nil, nil)
}

func (h *Helper) ResponseSuccessPayload(ctx *gin.Context, data interface{}) {
	h.Response(ctx, http.StatusOK, true, nil, data)
}

func (h *Helper) ResponseSuccessCreated(ctx *gin.Context, data interface{}) {
	h.Response(ctx, http.StatusCreated, true, nil, data)
}

func (h *Helper) ResponseNotFound(ctx *gin.Context) {
	h.Response(ctx, http.StatusNotFound, false, errors.New("not found"), nil)
}
