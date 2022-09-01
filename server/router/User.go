package router

import (
	"akexc.com/vueginlr-server/api"
	"akexc.com/vueginlr-server/middleware"
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(g *gin.RouterGroup)  {
	user := g.Group("user")
	user.Use(middleware.AuthMiddleware)
	{
		user.GET("/:id", api.GetUserById)
		//用户组  api/v1/user
		user.GET("/", api.GetAllUser)

		user.PUT("/:id", api.EditUser)

		user.DELETE("/:id",api.DelUser)
	}



}