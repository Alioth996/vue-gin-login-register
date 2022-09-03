package util

import (
	"time"

	"akexc.com/vueginlr-server/common"
	"akexc.com/vueginlr-server/model"
	"github.com/dgrijalva/jwt-go"
)

//定义秘钥
var jwtKey = []byte("akexc-jwt-4396")

//ReleaseToken 发放toke 有效期7天
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //token的有效期是七天
	claims := &common.Claims{
		UserId: user.ID,
		Role:   user.State, // 1 ==root,2 == admin, other == user
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //token的有效期
			IssuedAt:  time.Now().Unix(),     //token发放的时间
			Issuer:    "凌天akex",              //作者
			Subject:   "user token",          //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //根据前面自定义的jwt秘钥生成token

	if err != nil {
		//返回生成的错误
		return "", err
	}
	//返回成功生成的字符换
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *common.Claims, error) {
	claims := &common.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
