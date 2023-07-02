package routers

import (
	"github.com/gin-gonic/gin"
	"todo_backend/controller"
	"todo_backend/midddleware"
)

// FileRouter 文件路由
func FileRouter(engine *gin.Engine) {
	fr := engine.Group("file", midddleware.AccessLog(), midddleware.RequestAuth())
	{
		fc := controller.FileController{}
		fr.POST("", fc.CreateAttachment)
		fr.GET("/:filename", fc.DownloadAttachment)
		fr.DELETE("/:filename", fc.DelAttachment)
	}
}
