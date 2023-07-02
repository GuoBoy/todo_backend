package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo_backend/auth"
	"todo_backend/config"
	"todo_backend/controller"
	"todo_backend/models"
	"todo_backend/service"
)

// Login 登录
// @param 用户名、密码
// @return token
func Login(ctx *gin.Context) {
	var user models.AdminUserModel
	if err := controller.DecryptRequest[models.AdminUserModel](ctx, &user); err != nil {
		controller.ForbiddenResp(ctx, errors.New("form err"))
		return
	}
	adminCfg := config.Cfg.System.Admin
	for _, u := range adminCfg.Users {
		if u == user {
			token, err := auth.NewToken(1000000, user.Username)
			if err != nil {
				controller.ForbiddenResp(ctx, errors.New("generate token err"))
				return
			}
			controller.EncryptAnyResp(ctx, token)
			return
		}
	}
	controller.ForbiddenResp(ctx, errors.New("username or password is err"))
}

type Handles struct {
}

// GetLogs 获取日志
func (h Handles) GetLogs(ctx *gin.Context) {
	var query models.PaginationQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		controller.ForbiddenResp(ctx, err)
		return
	}
	sc := service.LogService{}
	res, err := sc.GetAllLogs(query)
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("加载日志失败"))
		return
	}
	controller.SuccessResp(ctx, res)
}

// GetAllUser 获取所有用户
func (Handles) GetAllUser(ctx *gin.Context) {
	userService := service.User{}
	users, err := userService.All()
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("获取失败"))
		return
	}
	controller.EncryptAnyResp(ctx, map[string]any{"data": users, "length": len(users)})
}

// ChangeUserEnabled 更改用户禁用状态
func (Handles) ChangeUserEnabled(ctx *gin.Context) {
	var form models.AdminUserEnabledForm
	if err := controller.DecryptRequest[models.AdminUserEnabledForm](ctx, &form); err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	userService := service.User{}
	if err := userService.SetUserEnabled(form); err != nil {
		controller.ForbiddenResp(ctx, errors.New("系统错误，修改失败"))
		return
	}
	controller.EncryptAnyResp(ctx, "修改成功")
}

// ResetUserPwd 重置用户密码
func ResetUserPwd(ctx *gin.Context) {
	tempId := ctx.Param("uid")
	id, err := strconv.Atoi(tempId)
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	userService := service.User{}
	if err = userService.UpdatePwdToDefault(id); err != nil {
		controller.ForbiddenResp(ctx, errors.New("修改失败"))
		return
	}
	controller.SuccessResp(ctx, "默认密码为"+config.Cfg.DefaultResetPwd)
}

func (Handles) DelUser(ctx *gin.Context) {
	tempId := ctx.Param("uid")
	id, err := strconv.Atoi(tempId)
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	userService := service.User{}
	if err = userService.DeleteUserByID(id); err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}
