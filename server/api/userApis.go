package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"akexc.com/vueginlr-server/model"
	"akexc.com/vueginlr-server/util"
	"github.com/gin-gonic/gin"
)

// var db = model.GetDb()

type userDto struct {
	Username string
	Nickname string
	State    int
	Phone    string
	Email    string
	Address  string
}

func Register(ctx *gin.Context) {
	var user model.User
	db := model.GetDb()
	_ = ctx.ShouldBindJSON(&user)

	if !checkUserExist(user.Username) {
		user.Password = util.GetHashPwd(user.Password)
		tx := db.Create(&user)
		if tx.Error != nil {
			log.Fatal("err", tx.Error)
		}
		util.Response(ctx, http.StatusCreated, 2001, nil, "注册成功")
	} else {
		util.Response(ctx, http.StatusBadRequest, 400, nil, "用户已存在")
	}

}

// Login 邮箱登录,账户密码登录,手机验证码登录,扫码登录
func Login(ctx *gin.Context) {
	var user model.User
	db := model.GetDb()
	_ = ctx.ShouldBindJSON(&user)
	password := user.Password
	if checkUserExist(user.Username) {
		db.Where("username = ? ", user.Username).First(&user)
		fmt.Println(password, user.Password)
		if util.ComparePwd(user.Password, password) {
			token, _ := util.ReleaseToken(user)
			util.Response(ctx, http.StatusOK, 2000, token, "登录成功")
		} else {
			util.Response(ctx, http.StatusBadRequest, 4000, nil, "用户名或密码错误")

		}

	} else {
		util.Response(ctx, http.StatusBadRequest, 4000, nil, "用户不存在")
	}
}

// GetUserInfo 用户获取个人信息
func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	util.Success(ctx, user, "用户信息获取成功")
}

// EditUser  管理员修改用户信息
func EditUser(ctx *gin.Context) {
	var user userDto
	err := ctx.Bind(&user)
	if err != nil {
		log.Fatal(err)
		return
	}
	id := ctx.Param("id")
	db := model.GetDb()

	if !checkUserExistById(id) {
		ctx.JSON(400, gin.H{
			"code": 4000,
			"msg":  "用户不存在",
		})
		return
	}
	// 根据 `struct` 更新属性，只会更新非零值的字段
	db.Model(model.User{}).Where("id = ?", id).Updates(user)
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	ctx.JSON(200, gin.H{
		"code": "2000",
		"msg":  "用户信息更新成功",
		"data": user,
	})

}

// DelUser root 删除用户
func DelUser(ctx *gin.Context) {
	id := ctx.Param("id")
	db := model.GetDb()

	if !checkUserExistById(id) {
		ctx.JSON(400, gin.H{
			"code": 4000,
			"msg":  "用户不存在",
		})
		return
	}
	db.Where("id = ?", id).Delete(&model.User{})
	ctx.JSON(200, gin.H{
		"code": "2000",
		"msg":  "用户已删除",
		"data": nil,
	})

}

// GetUserById 管理员获取用户信息
func GetUserById(ctx *gin.Context) {
	var user model.User

	id := ctx.Param("id")
	db := model.GetDb()

	if !checkUserExistById(id) {
		ctx.JSON(400, gin.H{
			"code": 4000,
			"msg":  "用户不存在",
		})
		return
	}
	db.First(&user, id)
	ctx.JSON(200, gin.H{
		"code": "2000",
		"msg":  "用户信息获取成成功",
		"data": user,
	})

}

// GetAllUser 管理员获取用户列表
func GetAllUser(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pagNum", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "5"))
	var users []model.User
	db := model.GetDb()
	// 返回分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (pageNum - 1) * pageSize
	// 查询所有的user
	result := db.Offset(offset).Limit(pageSize).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		ctx.JSON(200, gin.H{
			"code":  "2000",
			"msg":   "无数据",
			"total": 0,
			"data":  nil,
		})
		return
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(pageSize).Find(&users)
	for _, v := range users {
		userItem := map[string]interface{}{
			"id":       v.ID,
			"nickname": v.Nickname,
			"address":  v.Address,
			"state":    v.State,
			"email":    v.Email,
			"username": v.Username,
			"mobile":   v.Phone,
		}
		userList = append(userList, userItem)
	}
	ctx.JSON(200, gin.H{
		"code":  "2000",
		"msg":   "用户列表获取成功",
		"total": total,
		"data":  userList,
	})
}

// checkUserExist 根据用户名查询用户是否存在
func checkUserExist(name string) bool {
	var user model.User
	db := model.GetDb()
	db.Select("id").Where("username = ?", name).First(&user)

	if user.ID > 0 {
		return true
	} else {
		return false

	}

}

func checkUserExistById(id string) bool {
	var user model.User
	db := model.GetDb()
	db.Where("id = ?", id).First(&user)

	if user.ID > 0 {
		return true
	} else {
		return false

	}

}
