package common

import "github.com/gin-gonic/gin"

// AutoRoites 自动根据表名生成对应的CRUD路由
func AutoRoites(r *gin.Engine,rPath string,)  {
	r.GET("", func(context *gin.Context) {
		
	})
	
	r.POST("", func(context *gin.Context) {
		
	})
	
	r.PUT("", func(context *gin.Context) {
		
	})
	r.DELETE("", func(context *gin.Context) {
		
	})
}
