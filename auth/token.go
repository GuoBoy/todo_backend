package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"todo_backend/config"
	"todo_backend/models"
)

type CustomClaims struct {
	models.UserTokenInfo
	jwt.RegisteredClaims
}

// NewToken 生成token
func NewToken(uid uint, username string) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	claims.Uid = uid
	claims.Username = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.Cfg.JwtSecretBytes())
	if err != nil {
		fmt.Println("生成token失败", err)
		return "", err
	}
	return tokenString, nil
}

// VerifyByToken 验证token
func VerifyByToken(tokenString string) (models.UserTokenInfo, bool) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) { return config.Cfg.JwtSecretBytes(), nil })
	if err != nil {
		return models.UserTokenInfo{}, false
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserTokenInfo, true
	}
	return models.UserTokenInfo{}, false
}
