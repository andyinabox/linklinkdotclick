package ginhelper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUint(ctx *gin.Context, key string) uint {
	return uint(ctx.GetInt(key))
}

func GetParamUint(ctx *gin.Context, key string) (uint, error) {
	str := ctx.Param(key)
	v, err := strconv.Atoi(str)
	return uint(v), err
}

func GetPostFormUint(ctx *gin.Context, key string) (uint, error) {
	str := ctx.PostForm(key)
	v, err := strconv.Atoi(str)
	return uint(v), err
}

func GetQueryUint(ctx *gin.Context, key string) (uint, error) {
	str, _ := ctx.GetQuery(key)
	v, err := strconv.Atoi(str)
	return uint(v), err
}

func GetQueryBool(ctx *gin.Context, key string) bool {
	result, _ := ctx.GetQuery(key)
	if result == "true" || result == "1" {
		return true
	}
	return false
}
