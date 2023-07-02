package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo_backend/apps/admin"
	"todo_backend/apps/booknote_app"
	"todo_backend/apps/quadrant_todo"
	"todo_backend/apps/todo_app"
	"todo_backend/config"
	"todo_backend/controller"
	"todo_backend/env"
	"todo_backend/midddleware"
)

// 注册路由
func registerGroup(r *gin.Engine, routerFunc ...func(engine *gin.Engine)) {
	for _, fun := range routerFunc {
		fun(r)
	}
}

// Run 启动路由
func Run() {
	gin.SetMode(gin.ReleaseMode)
	if env.Env.DevelopmentEnv {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.Use(midddleware.CorsMiddleware())
	r.Any("/", midddleware.AccessLog(), func(context *gin.Context) {
		context.JSON(200, gin.H{"code": 200, "msg": "pong"})
	})
	r.GET("/todo-ws", midddleware.AccessLog(), controller.WsHandle)
	r.GET("/version", midddleware.AccessLog(), controller.GetVersion)
	r.POST("/feedback", midddleware.AccessLog(), midddleware.RequestAuth(), controller.OnFeedBack)
	registerGroup(r, LR, todo_app.GroupRouter, todo_app.ItemRouter, VersionRoute, RecycleBin, admin.Admin,
		booknote_app.BookNoteRouter, FileRouter, quadrant_todo.QuadrantRouter)
	r.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	log.Fatal(r.Run(fmt.Sprintf(":%d", config.Cfg.ServerPort)))
}
