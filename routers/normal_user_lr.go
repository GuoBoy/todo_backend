package routers

import (
	"github.com/gin-gonic/gin"
	"todo_backend/controller"
	"todo_backend/midddleware"
)

// LR 登录注册
func LR(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	authed := r.Group("", midddleware.RequestAuth())
	{
		authed.POST("/changepwd", controller.ChangePassword)
		authed.GET("/verify", func(context *gin.Context) {
			controller.SuccessResp(context, "ok")
		})
	}
}
