package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo_backend/auth"
	"todo_backend/models"
	"todo_backend/service"
)

// Login 登录
// @param 用户名、密码
// @return token
func Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ForbiddenResp(ctx, errors.New("form err"))
		return
	}
	userService := service.User{D: &user}
	if err := userService.VerifyByUaPW(); err != nil {
		ForbiddenResp(ctx, errors.New("username or password is err"))
		return
	}
	token, err := auth.NewToken(user.ID, user.Username)
	if err != nil {
		ForbiddenResp(ctx, errors.New("generate token err"))
		return
	}
	//ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "uid": verifiedUser.ID, "last_update_time": verifiedUser.UpdateTime})
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "token": token})
}

// Register 注册
// @param 用户名、密码
// @return token
func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ForbiddenResp(ctx, errors.New("form err"))
		return
	}
	userService := service.User{D: &user}
	if err := userService.Create(); err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	token, err := auth.NewToken(user.ID, user.Username)
	if err != nil {
		ForbiddenResp(ctx, errors.New("generate token err"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "token": token})
}

// CheckUpdate
// 检查用户本地是否更新
func CheckUpdate(ctx *gin.Context) {
	/**
	1. 获取update_time
	2. 对比本地和远程大小
	3. 发送是否需要更新
	*/
	uid := ctx.MustGet("uid").(uint)
	tempUt := ctx.Param("ut")
	ut, err := strconv.ParseInt(tempUt, 10, 64)
	if err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	userService := service.User{}
	ok, err := userService.CheckUpdateTime(uid, ut)
	if err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	// 响应是否需要更新
	switch ok {
	case true:
		ctx.JSON(http.StatusOK, gin.H{"code": 302, "msg": "need update", "updated_at": userService.D.UpdateTime})
		break
	default:
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "已是最新"})
	}
}

// ChangePassword 修改密码
func ChangePassword(ctx *gin.Context) {
	var form models.PwdForm
	if err := DecryptRequest[models.PwdForm](ctx, &form); err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	userService := service.User{}
	if err := userService.UpdatePwd(ctx.MustGet("uid").(uint), form); err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	EncryptResp(ctx, EncryptToForm(gin.H{"code": 200}))
}

// OnFeedBack 收到反馈信息
func OnFeedBack(ctx *gin.Context) {
	var form models.FeedBackModel
	if err := ctx.ShouldBind(&form); err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	form.Uid = ctx.MustGet("uid").(uint)
	fbs := service.FeedBackService{}
	if err := fbs.StoreFeedBack(&form); err != nil {
		ForbiddenResp(ctx, err)
		return
	}
	SuccessResp(ctx, "thanks to give feedback!")
}
