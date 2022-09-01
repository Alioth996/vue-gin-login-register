package router

import (
	"net/http"

	"akexc.com/vueginlr-server/api"
	"akexc.com/vueginlr-server/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("debug")
	r := gin.Default()
	// 使用Cors解决跨域
	r.Use(util.Cors())
	// 静态资源服务
	r.Static("/assets", "./assets")

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "pong",
			"ip":   ctx.ClientIP(),
		})
	})
	r.POST("api/v1/login", api.Login)
	r.POST("api/v1/register", api.Register)

	return r
}
