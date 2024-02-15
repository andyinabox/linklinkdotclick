package handlerhelper

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Helper struct {
	sc app.ServiceContainer
}

func New(sc app.ServiceContainer) *Helper {
	return &Helper{sc}
}

type JSONResponseBody struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func (h *Helper) ResponseJSON(ctx *gin.Context, code int, success bool, err error, data interface{}) {

	response := &JSONResponseBody{
		Success: success,
		Payload: data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	ctx.JSON(code, response)
}

func (h *Helper) ErrorResponse(ctx *gin.Context, code int, err error) {
	h.ResponseJSON(ctx, code, false, err, nil)
}

func (h *Helper) SuccessResponse(ctx *gin.Context) {
	h.ResponseJSON(ctx, http.StatusOK, true, nil, nil)
}

func (h *Helper) SuccessResponseJSON(ctx *gin.Context, data interface{}) {
	h.ResponseJSON(ctx, http.StatusOK, true, nil, data)
}

func (h *Helper) CreatedResponseJSON(ctx *gin.Context, data interface{}) {
	h.ResponseJSON(ctx, http.StatusCreated, true, nil, data)
}

func (h *Helper) NotFoundResponse(ctx *gin.Context) {
	h.ResponseJSON(ctx, http.StatusNotFound, false, errors.New("not found"), nil)
}

func (h *Helper) GetIdParam(ctx *gin.Context) (uint, error) {
	str := ctx.Param("id")
	id, err := strconv.Atoi(str)
	return uint(id), err
}

func (h *Helper) GetUserIdFromSession(ctx *gin.Context) (id uint, isDefaultUser bool, err error) {
	session := sessions.Default(ctx)
	value := session.Get("user")
	if value == nil {
		return h.sc.UserService().GetDefaultUserId(), true, nil
	}
	var ok bool
	id, ok = value.(uint)
	if !ok {
		err = app.ErrServerError
		return
	}
	return
}

func (h *Helper) GetUserFromSession(ctx *gin.Context) (*app.User, bool, error) {
	id, isDefaultUser, err := h.GetUserIdFromSession(ctx)
	if err != nil {
		return nil, false, err
	}
	user, err := h.sc.UserService().FetchUser(id)
	return user, isDefaultUser, err
}
