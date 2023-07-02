package admin

import (
	"github.com/gin-gonic/gin"
	"todo_backend/controller"
	"todo_backend/models"
	"todo_backend/service"
)

// VersionHandle 版本控制器
type VersionHandle struct {
}

// AddApp 创建app
func (VersionHandle) AddApp(ctx *gin.Context) {
	var app models.AppVersion
	if err := controller.DecryptRequest[models.AppVersion](ctx, &app); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	ser := service.AppService{}
	id, err := ser.CreateApp(app)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, id)
}

// GetAllApp 获取所有app
func (VersionHandle) GetAllApp(ctx *gin.Context) {
	ser := service.AppService{}
	res, err := ser.GetAllApp()
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, res)
}

// AddHistory 添加app历史
func (VersionHandle) AddHistory(ctx *gin.Context) {
	var his models.AppHistoryVersion
	if err := controller.DecryptRequest[models.AppHistoryVersion](ctx, &his); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	ser := service.AppService{}
	id, err := ser.CreateAppHistory(his)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, id)
}

// GetAppHis 获取app更新历史
func (VersionHandle) GetAppHis(ctx *gin.Context) {
	id := ctx.Param("app_id")
	ser := service.AppService{}
	res, err := ser.GetAppHistoryByID(id)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, res)
}
