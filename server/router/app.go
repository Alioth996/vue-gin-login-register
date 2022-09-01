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

	// test connect
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "pong",
			"ip":   ctx.ClientIP(),
		})
	})



	//用户组
	v1 := r.Group("api/v1")
	{
		// 特殊接口,不需要token严重
		v1.POST("/login", api.Login)
		v1.POST("/register", api.Register)
		GetUserRoutes(v1)
	}



	return r
}
