package midddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_backend/auth"
)

// RequestAuth 请求验证
func RequestAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		user, ok := auth.VerifyByToken(tokenString)
		if ok {
			ctx.Set("uid", user.Uid)
			ctx.Set("userinfo", user)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
