package todo_app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_backend/controller"
	"todo_backend/models"
)

type GroupController struct {
}

// Add 新增分组
func (GroupController) Add(ctx *gin.Context) {
	var form models.GroupModel
	if err := ctx.ShouldBind(&form); err != nil {
		controller.ForbiddenResp(ctx, errors.New("新增分组失败"))
		return
	}
	form.UserID = ctx.MustGet("uid").(uint)
	groupService := GroupService{D: &form}
	ut, err := groupService.Create()
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("新增分组失败"))
		return
	}
	controller.SuccessResp(ctx, gin.H{"id": groupService.D.ID, "update_time": ut})
}

// Del 删除分组
func (GroupController) Del(ctx *gin.Context) {
	gid := ctx.Param("gid")
	if err := CheckGroupByID(gid); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	uid := ctx.MustGet("uid").(uint)
	groupService := GroupService{}
	ut, err := groupService.Delete(gid, uid)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, ut)
	//ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "update_time": ut})
}

// FDel 强制删除
func (GroupController) FDel(ctx *gin.Context) {
	gid := ctx.Param("gid")
	uid := ctx.MustGet("uid").(uint)
	groupService := GroupService{}
	ut, err := groupService.ForceDelete(gid, uid)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, ut)
	//ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "update_time": ut})
}

// UpdateName 更新分组名称
func (GroupController) UpdateName(ctx *gin.Context) {
	gid := ctx.Param("gid")
	if err := CheckGroupByID(gid); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	// 解析数据
	var form models.GroupModel
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 403, "msg": "表单错误"})
		return
	}
	uid := ctx.MustGet("uid").(uint)
	groupService := GroupService{}
	ut, err := groupService.Update(uid, gid, form.Name)
	if err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	controller.SuccessResp(ctx, ut)
}

// GetAll 获取所有分组
func (GroupController) GetAll(ctx *gin.Context) {
	groupService := GroupService{}
	ms, err := groupService.All(ctx.MustGet("uid").(uint))
	if err != nil {
		controller.EncryptResp(ctx, controller.EncryptToForm(gin.H{"code": 403, "msg": err.Error()}))
		return
	}
	controller.EncryptAnyResp(ctx, map[string]any{"data": ms, "length": len(ms)})
}
