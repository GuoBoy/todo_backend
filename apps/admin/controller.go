package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"todo_backend/apps/todo_app"
	"todo_backend/controller"
	"todo_backend/models"
)

type Controller struct {
}

// GetAllTodoItems 获取所有item
func (Controller) GetAllTodoItems(ctx *gin.Context) {
	items, err := AllItems()
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("失败"))
		return
	}
	controller.EncryptAnyResp(ctx, items)
}

// ChangeItemDetail 修改item
func (Controller) ChangeItemDetail(ctx *gin.Context) {
	var form models.ChangeItemDetailForm
	if err := controller.DecryptRequest[models.ChangeItemDetailForm](ctx, &form); err != nil {
		controller.ForbiddenResp(ctx, errors.New("params is err"))
		return
	}
	if err := UpdateItemDetail(form.ID, form.Detail); err != nil {
		controller.ForbiddenResp(ctx, errors.New("system err"))
		return
	}
	controller.EncryptAnyResp(ctx, "success")
}

func (Controller) DelItem(ctx *gin.Context) {
	tempId := ctx.Param("id")
	id, err := strconv.Atoi(tempId)
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("params is err"))
		return
	}
	if err = DelItem(id); err != nil {
		controller.ForbiddenResp(ctx, errors.New("system is err"))
		return
	}
	controller.SuccessResp(ctx, "success")
}

// ChangeChecked 修改item状态
func (Controller) ChangeChecked(ctx *gin.Context) {
	uid := ctx.MustGet("uid").(uint)
	// 解析item id
	tempTid := ctx.Param("tid")
	tid, err := strconv.Atoi(tempTid)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	tempChecked := ctx.Param("checked")
	checked, _ := strconv.ParseBool(tempChecked)
	ser := todo_app.ItemService{}
	if _, err = ser.UpdateChecked(int(uid), tid, checked); err != nil {
		controller.ForbiddenResp(ctx, errors.New("system is err"))
		return
	}
	controller.SuccessResp(ctx, "success")
}
