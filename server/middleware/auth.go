package middleware

import (
	"fmt"
	"strings"

	"akexc.com/vueginlr-server/model"
	"akexc.com/vueginlr-server/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context)  {
		// 获取 authorization header：获取前端传过来的信息的
		tokenString := c.GetHeader("Authorization")
		//fmt.Println(tokenString)
		//验证前端传过来的token格式，不为空，开头为Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401,gin.H{
				"code":4001,
				"msg":"token格式错误",
			})
			c.Abort()  // 结束调用,不然会继续调用接口
			return
		}

		//验证通过，提取有效部分（除去Bearer)
		tokenString = tokenString[7:] //截取字符
		token, claims, err := util.ParseToken(tokenString)
		//解析失败||解析后的token无效
		if err != nil || !token.Valid {
			c.JSON(401,gin.H{
				"code":4001,
				"msg":"无效token",
			})
			c.Abort()
			return
		}

		//token通过验证, 获取claims中的UserID
		userId := claims.UserId
		var user model.User
		db := model.GetDb()
		//查询数据库
		db.First(&user, userId)

		fmt.Println(userId)
		// 验证用户是否存在
		if user.ID != 1 {
			c.JSON(401, gin.H{
				"code":4001,
				"msg":"权限不足,非管理员",
			})
			c.Abort()
			return
		}

		// token有效,存入路由中
		c.Set("user",claims)
}
