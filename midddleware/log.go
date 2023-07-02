package midddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"todo_backend/models"
	"todo_backend/service"
)

// AccessLog 日志中间件
func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		// Process request
		ctx.Next()
		var ut models.UserTokenInfo
		if user, ok := ctx.Get("userinfo"); ok {
			ut = user.(models.UserTokenInfo)
		}
		// 计算时长
		current := time.Now()
		logm := models.AccessLogModel{
			CreatedAt:  current,
			ClientIP:   ctx.ClientIP(),
			UserAgent:  ctx.Request.UserAgent(),
			Method:     ctx.Request.Method,
			Path:       path,
			StatusCode: ctx.Writer.Status(),
			UID:        ut.Uid,
			Username:   ut.Username,
			Latency:    current.Sub(start),
		}
		if raw != "" {
			logm.Path = path + "?" + raw
		}
		//fmt.Printf("%d %s |%d| %s | %s| %s '%s'\n", logm.UID, logm.CreatedAt.Format(config.TimeLayout), logm.StatusCode, logm.Latency, logm.ClientIP, logm.Method, logm.Path)
		if err := service.AddLog(logm); err != nil {
			fmt.Println(err)
		}
	}
}
