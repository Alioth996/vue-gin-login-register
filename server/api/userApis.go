package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"akexc.com/vueginlr-server/model"
	"akexc.com/vueginlr-server/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// var db = model.GetDb()

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
			token, _ := CreateToken(&user, "secret64h")
			util.Response(ctx, http.StatusOK, 2000, token, "登录成功")
		} else {
			util.Response(ctx, http.StatusBadRequest, 4000, nil, "用户名或密码错误")

		}

	} else {
		util.Response(ctx, http.StatusBadRequest, 4000, nil, "用户不存在")
	}
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

func CreateToken(info *model.User, secret string) (tokenString string, err error) {
	//传入指定的签名方法和payload信息,创建Token对象
	//库中内置了好几种签名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       info.ID,
		"username": info.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	//根据密码生成token字符串
	tokenString, err = token.SignedString([]byte(secret))
	return tokenString, err
}
