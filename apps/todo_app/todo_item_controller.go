package todo_app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"todo_backend/controller"
	"todo_backend/models"
	"todo_backend/service"
	"todo_backend/utils"
)

type ItemController struct {
}

// Add 新增todo item
func (ItemController) Add(ctx *gin.Context) {
	// 解析表单
	var form models.ItemForm
	if err := controller.DecryptRequest[models.ItemForm](ctx, &form); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	uid := ctx.MustGet("uid").(uint)
	itemService := ItemService{D: &models.ItemModel{
		Detail:  form.Detail,
		GroupId: form.GroupID,
		Checked: false,
		UserID:  uid,
	}, UserID: uid}
	ut, err := itemService.Create()
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("新增todo失败"))
		return
	}
	item := itemService.D
	data := map[string]any{"id": item.ID, "hashcode": item.Hashcode, "updated_time": ut, "updated": utils.FormatTime2Standard(item.UpdatedAt)}
	controller.EncryptAnyResp(ctx, data)
}

// Del 删除分组控制器
func (ItemController) Del(ctx *gin.Context) {
	tid := ctx.Param("tid")
	uid := ctx.MustGet("uid").(uint)
	itemService := ItemService{}
	ut, err := itemService.Delete(tid, uid)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, ut)
}

func (ItemController) FDel(ctx *gin.Context) {
	tid := ctx.Param("tid")
	uid := ctx.MustGet("uid").(uint)
	itemService := ItemService{}
	ut, err := itemService.ForceDelete(tid, uid)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, ut)
}

// Update 更新item
func (ItemController) Update(ctx *gin.Context) {
	uid := ctx.MustGet("uid").(uint)
	// 解析item id
	tid := ctx.Param("tid")
	// 解析数据
	var form models.ItemForm
	if err := controller.DecryptRequest[models.ItemForm](ctx, &form); err != nil && form.Hashcode != "" {
		controller.ForbiddenResp(ctx, err)
		return
	}
	itemService := ItemService{}
	ut, err := itemService.Update(uid, tid, form.Hashcode, form.Detail)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	item := itemService.D
	data := map[string]any{"updated_at": ut, "hashcode": item.Hashcode, "updated": utils.FormatTime2Standard(item.UpdatedAt)}
	controller.EncryptAnyResp(ctx, data)
}

// UpdateState 更新状态
func (ItemController) UpdateState(ctx *gin.Context) {
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
	itemService := ItemService{}
	ut, err := itemService.UpdateChecked(int(uid), tid, checked)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	item := itemService.D
	controller.SuccessResp(ctx, gin.H{"updated_at": ut, "updated": utils.FormatTime2Standard(item.UpdatedAt)})
}

func (ItemController) GetAll(ctx *gin.Context) {
	itemService := ItemService{}
	ms, err := itemService.GetAllByUid(ctx.MustGet("uid").(uint))
	if err != nil {
		controller.EncryptResp(ctx, controller.EncryptToForm(gin.H{"code": 403, "data": err.Error()}))
		return
	}
	controller.EncryptResp(ctx, controller.EncryptToForm(map[string]any{"data": ms, "length": len(ms)}))
}

// GetAllByGroupID 根据分组id获取所有item
func (ItemController) GetAllByGroupID(ctx *gin.Context) {
	gid := ctx.Param("gid")
	itemService := ItemService{}
	ms, err := itemService.GetAllByGroupID(gid)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	ut, err := service.GetUserUpdateTime(ctx.MustGet("uid").(uint))
	if err != nil {
		controller.ForbiddenResp(ctx, err)
	}
	controller.EncryptAnyResp(ctx, map[string]any{"data": ms, "length": len(ms), "updated_at": ut})
}
