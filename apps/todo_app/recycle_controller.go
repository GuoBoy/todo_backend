package todo_app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"todo_backend/controller"
	"todo_backend/service"
)

// RecycleBinController 回收站控制器
type RecycleBinController struct {
}

// GetAll 加载回收站数据
func (RecycleBinController) GetAll(ctx *gin.Context) {
	recycleService := service.RecycleService{}
	if err := recycleService.All(ctx.MustGet("uid").(uint)); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, gin.H{"items": recycleService.Items, "groups": recycleService.Groups})
}

// Recover 恢复回收站数据，item or group
func (RecycleBinController) Recover(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := ctx.MustGet("uid").(uint)
	recycleService := service.RecycleService{}
	var (
		err error
		dt  int64
	)
	rType := ctx.Query("type")
	switch rType {
	case "group":
		dt, err = recycleService.GroupRecoverByID(id, uid)
		break
	case "item":
		dt, err = recycleService.ItemRecoverByID(id, uid)
		break
	default:
		err = errors.New("参数错误")
	}
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, dt)
}
