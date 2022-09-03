package router

import (
	"akexc.com/vueginlr-server/api"
	"akexc.com/vueginlr-server/common"
	"akexc.com/vueginlr-server/middleware"
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(g *gin.RouterGroup) {
	////用户组  api/v1/user
	user := g.Group("user")
	{
		//	需要登录
		userAuth := user.Group("")
		userAuth.Use(middleware.AuthMiddleware(common.User))
		{
			userAuth.GET("/info", api.GetUserInfo)
			userAuth.PUT("/:id", api.EditUser) // 修改用户个人信息
		}
		//管理员
		adminAuth := user.Group("")
		adminAuth.Use(middleware.AuthMiddleware(common.Admin))
		{
			adminAuth.GET("/:id", api.GetUserById) // 通过id获取用户信息
			adminAuth.GET("/", api.GetAllUser)     // 获取用户列表
		}
		// 超级管理员
		rootAuth := user.Group("")
		rootAuth.Use(middleware.AuthMiddleware(common.Root))
		{
			rootAuth.DELETE("/:id", api.DelUser) // 删除单个用户
		}

		//
	}

}
