package quadrant_todo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"todo_backend/controller"
	"todo_backend/models"
	"todo_backend/utils"
)

type Handle struct {
}

// Add 新增todo item
func (Handle) Add(ctx *gin.Context) {
	// 解析表单
	var form models.QTodoModel
	if err := controller.DecryptRequest[models.QTodoModel](ctx, &form); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	form.UserID = ctx.MustGet("uid").(uint)
	serv := Service{}
	if err := serv.Create(&form); err != nil {
		controller.ForbiddenResp(ctx, errors.New("新增todo失败"))
		return
	}
	data := map[string]any{"id": form.ID, "hashcode": form.Hashcode, "updated": utils.FormatTime2Standard(form.UpdatedAt)}
	controller.EncryptAnyResp(ctx, data)
}

// Del 删除item控制器
func (Handle) Del(ctx *gin.Context) {
	tid := ctx.Param("tid")
	uid := ctx.MustGet("uid").(uint)
	sev := Service{}
	if err := sev.Delete(tid, uid); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, "success")
}

// GetAll 获取所有 q item
func (Handle) GetAll(ctx *gin.Context) {
	tempType := ctx.Param("types")
	types, _ := strconv.ParseUint(tempType, 10, 32)
	itemService := Service{}
	ms, err := itemService.GetAllByQuadrantType(ctx.MustGet("uid").(uint), uint(types))
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, ms)
}

// UpdateDetail 更新详情
func (Handle) UpdateDetail(ctx *gin.Context) {
	id := ctx.Param("tid")
	type tempForm struct {
		Detail string `json:"detail" binding:"required"`
	}
	var form tempForm
	if err := controller.DecryptRequest[tempForm](ctx, &form); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	itemService := Service{}
	item, err := itemService.UpdateDetail(id, form.Detail)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.EncryptAnyResp(ctx, gin.H{"hashcode": item.Hashcode, "updated_at": item.UpdatedAt})
}
