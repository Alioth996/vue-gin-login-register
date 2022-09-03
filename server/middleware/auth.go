package middleware

import (
	"strings"

	"akexc.com/vueginlr-server/model"
	"akexc.com/vueginlr-server/util"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户鉴权 传入 role
func AuthMiddleware(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 authorization header：获取前端传过来的信息的
		tokenString := c.GetHeader("Authorization")
		//fmt.Println(tokenString)
		//验证前端传过来的token格式，不为空，开头为Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{
				"code": 4001,
				"msg":  "token格式错误",
			})
			c.Abort() // 抛弃请求
			return
		}

		//验证通过，提取有效部分（除去Bearer)
		tokenString = tokenString[7:] //截取字符
		token, claims, err := util.ParseToken(tokenString)

		//fmt.Println(claims.Role,role)
		//解析失败||解析后的token无效
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{
				"code": 4001,
				"msg":  "无效token",
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

		if claims.Role < role {
			c.JSON(401, gin.H{
				"code": 4001,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		//将用户信息写入上下文
		c.Set("user", user)
		//将用户id和权限写入上下文
		c.Set("uid", userId)
		c.Set("role", claims.Role)
		c.Next()
	}

}
