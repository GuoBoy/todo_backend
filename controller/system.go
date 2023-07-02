package controller

import (
	"github.com/gin-gonic/gin"
	"todo_backend/service"
)

// 版本检查参数
type versionCheckQuery struct {
	Platform string `form:"platform" binding:"required"`
	Value    int64  `form:"value"  binding:"required"`
}

// GetVersion 检查是否需要更新，并返回版本信息
func GetVersion(ctx *gin.Context) {
	var query versionCheckQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	val, name, err := service.GetVersion(query.Platform)
	if err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	data := gin.H{"update": false}
	if val > query.Value {
		data = gin.H{"update": true, "value": val, "name": name}
	}
	SuccessResp(ctx, data)
}
