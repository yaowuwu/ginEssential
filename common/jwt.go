package common

import (
	"github.com/dgrijalva/jwt-go"
	"hello/ginessential/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct{
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User)(string, error){
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "hello",
			Subject: "user.token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if(err != nil){
		return "", err
	}
	return tokenString, nil
}

//命令行解析 token 三部分 但我试了找不到-D参数
//echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 | base64 -D

func ParseToken(tokenString string)(*jwt.Token, *Claims, error){
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(i interface{}, err error){
		return jwtKey, nil
	})
	return  token, claims, err
}