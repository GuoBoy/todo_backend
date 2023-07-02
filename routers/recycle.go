package routers

import (
	"github.com/gin-gonic/gin"
	"todo_backend/apps/todo_app"
	"todo_backend/midddleware"
)

// RecycleBin 回收站路由
func RecycleBin(r *gin.Engine) {
	b := r.Group("recycle", midddleware.RequestAuth(), midddleware.AccessLog())
	{
		recycleController := todo_app.RecycleBinController{}
		b.GET("", recycleController.GetAll)
		b.PUT("/:id", recycleController.Recover)
	}
}
