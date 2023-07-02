package quadrant_todo

import (
	"github.com/gin-gonic/gin"
	"todo_backend/midddleware"
)

// QuadrantRouter todo分组路由表
func QuadrantRouter(r *gin.Engine) {
	qr := r.Group("/quadrant", midddleware.AccessLog(), midddleware.RequestAuth())
	handle := Handle{}
	{
		qr.POST("", handle.Add)               // 添加 encrypt
		qr.DELETE("/:tid", handle.Del)        //删除
		qr.POST("/:tid", handle.UpdateDetail) // 更新
		qr.GET("/:types", handle.GetAll)
	}
}
