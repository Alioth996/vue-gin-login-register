package common

import "github.com/dgrijalva/jwt-go"

//jwt payload 数据

type Claims struct {
	UserId uint
	Role   int
	jwt.StandardClaims
}
