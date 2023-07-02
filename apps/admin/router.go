package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_backend/config"
	"todo_backend/front-admin"
	"todo_backend/midddleware"
)

func Admin(r *gin.Engine) {
	adminCfg := config.Cfg.System.Admin
	// 静态资源
	r.Any(adminCfg.Uri, func(ctx *gin.Context) {
		ctx.Writer.Write(front_admin.Index)
	})
	r.Any("/assets/*filepath", func(ctx *gin.Context) {
		serv := http.FileServer(http.FS(front_admin.Assets))
		serv.ServeHTTP(ctx.Writer, ctx.Request)
	})
	// robots
	r.Any("/robots.txt", func(ctx *gin.Context) {
		ctx.Writer.Write(front_admin.RobotsFile)
	})
	// api接口
	admin := r.Group(adminCfg.Uri, midddleware.AccessLog())
	{
		admin.POST("/login", Login)
		// 用户相关api
		userMange := admin.Group("/user", midddleware.RequestAuth())
		{
			userCont := Handles{}
			userMange.GET("", userCont.GetAllUser)
			userMange.POST("/state", userCont.ChangeUserEnabled)
			//userMange.DELETE("/:uid", userCont.DelUser)
			userMange.PUT("/pwd/:uid", ResetUserPwd)
		}
		//item相关api
		itemG := admin.Group("item", midddleware.RequestAuth())
		{
			adminController := Controller{}
			itemG.GET("", adminController.GetAllTodoItems)
			itemG.POST("", adminController.ChangeItemDetail)
			itemG.DELETE("/:id", adminController.DelItem)
			itemG.PUT("/:tid/:checked", adminController.ChangeChecked)
		}
		// log
		logG := admin.Group("log", midddleware.RequestAuth())
		{
			handles := Handles{}
			logG.GET("", handles.GetLogs)
		}
		// app 版本管理
		ver := admin.Group("appVersion", midddleware.RequestAuth())
		{
			verHand := VersionHandle{}
			ver.POST("/", verHand.AddApp)
			ver.GET("/", verHand.GetAllApp)
			ver.POST("/history", verHand.AddHistory)
			ver.GET("/:app_id", verHand.GetAppHis)
		}
	}
}
