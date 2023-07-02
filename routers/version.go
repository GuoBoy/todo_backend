package routers

import (
	"github.com/gin-gonic/gin"
	"todo_backend/version"
)

func VersionRoute(r *gin.Engine) {
	v := r.Group("/version")
	{
		v.GET("/pc", func(context *gin.Context) {
			context.JSON(200, gin.H{"code": 200, "v": version.GetV(version.TauriPc), "link": version.GetDownloadLink()})
		})
		v.GET("/android", func(context *gin.Context) {
			context.JSON(200, gin.H{"code": 200, "v": version.GetV(version.FlutterApp), "link": version.GetDownloadLink()})
		})
	}
}
