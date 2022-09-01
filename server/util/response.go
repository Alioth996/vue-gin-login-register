package util

import (
	"github.com/gin-gonic/gin"
)

// Response 响应请求
func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
