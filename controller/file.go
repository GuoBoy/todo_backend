package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"todo_backend/config"
	"todo_backend/models"
	"todo_backend/service"
	"todo_backend/utils"
)

// FileController 文件控制器
type FileController struct {
}

// CreateAttachment 插入附件
func (f *FileController) CreateAttachment(ctx *gin.Context) {
	user := ctx.MustGet("userinfo").(models.UserTokenInfo)
	//TODO 检查附件
	file, err := ctx.FormFile("file")
	if err != nil {
		ForbiddenResp(ctx, errors.New("获取文件失败"))
		return
	}
	// 检查文件类型
	if utils.IsForbiddenFile(file.Filename) {
		ForbiddenResp(ctx, errors.New("文件禁止上传"))
		return
	}
	fileStore := models.NewFileMode(file.Filename, file.Size, user.Uid)
	abs_path := path.Join(config.Cfg.UploadDir, fileStore.Filename)
	// 保存文件
	if err = ctx.SaveUploadedFile(file, abs_path); err != nil {
		ForbiddenResp(ctx, errors.New("获取文件失败"))
		return
	}
	// 保存信息
	ser := service.FileService{}
	if err = ser.Create(fileStore); err != nil {
		// 失败删除保存的文件
		os.Remove(abs_path)
		ForbiddenResp(ctx, errors.New("保存失败"))
		return
	}
	SuccessResp(ctx, fileStore.Filename)
}

// DelAttachment 删除附件
func (f *FileController) DelAttachment(ctx *gin.Context) {
	filename := ctx.Param("filename")
	ser := service.FileService{}
	if err := ser.Del(filename); err != nil {
		ForbiddenResp(ctx, errors.New("删除失败"))
		return
	}
	os.Remove(path.Join(config.Cfg.UploadDir, filename))
	SuccessResp(ctx, "ok")
}

// DownloadAttachment 下载附件
func (f *FileController) DownloadAttachment(ctx *gin.Context) {
	fn := ctx.Param("filename")
	if fn == "" {
		ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	ser := service.FileService{}
	ft, err := ser.GetInfo(fn)
	if err != nil {
		ForbiddenResp(ctx, errors.New("文件不存在或已删除"))
		return
	}
	ctx.FileAttachment(path.Join(config.Cfg.UploadDir, fn), ft.OriginName)
}
