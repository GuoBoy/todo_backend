package todo_app

import (
	"github.com/gin-gonic/gin"
	"todo_backend/controller"
	"todo_backend/midddleware"
)

// GroupRouter todo分组路由表
func GroupRouter(r *gin.Engine) {
	g := r.Group("/group", midddleware.RequestAuth(), midddleware.AccessLog())
	groupController := GroupController{}
	{
		g.POST("", groupController.Add)
		g.DELETE(":gid", groupController.Del)
		g.DELETE("/force/:gid", groupController.FDel)
		g.POST("/:gid", groupController.UpdateName)
		g.GET("", groupController.GetAll) // encryot
	}
}

// ItemRouter item路由
func ItemRouter(r *gin.Engine) {
	todo := r.Group("/todo", midddleware.RequestAuth(), midddleware.AccessLog())
	todoController := ItemController{}
	{
		todo.POST("", todoController.Add)               // 添加 encrypt
		todo.DELETE("/:tid", todoController.Del)        //删除
		todo.DELETE("/force/:tid", todoController.FDel) //强制删除
		todo.POST("/:tid", todoController.Update)       // 更新
		todo.POST("/:tid/:checked", todoController.UpdateState)
		todo.GET("/:gid", todoController.GetAllByGroupID)
		// 检查更新
		todo.GET("/checkUpdate/:ut", controller.CheckUpdate)
	}
}
