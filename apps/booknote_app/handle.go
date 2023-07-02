package booknote_app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"todo_backend/config"
	"todo_backend/controller"
	"todo_backend/models"
	"todo_backend/utils"
)

type BookController struct {
}

// GetAll 获取所有书籍
func (*BookController) GetAll(ctx *gin.Context) {
	ser := BookService{}
	res, ln, err := ser.GetAll()
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("获取失败"))
		return
	}
	controller.EncryptAnyResp(ctx, gin.H{"data": res, "length": ln})
}

// Create 新建书籍
func (b *BookController) Create(ctx *gin.Context) {
	var book models.BookNote
	if err := controller.DecryptRequest[models.BookNote](ctx, &book); err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数解析错误"))
		return
	}
	ser := BookService{D: &book}
	if err := ser.Add(); err != nil {
		controller.ForbiddenResp(ctx, errors.New("保存失败"))
		return
	}
	controller.EncryptAnyResp(ctx, gin.H{"id": book.ID, "created_at": utils.FormatTime2Standard(book.CreatedAt)})
}

// Del 删除书籍
func (b *BookController) Del(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	ser := BookService{}
	if ser.DelByID(id) != nil {
		controller.ForbiddenResp(ctx, errors.New("删除失败"))
		return
	}
	controller.SuccessResp(ctx, "删除成功")
}

// UpdateBookName 更新书籍名称
func (b *BookController) UpdateBookName(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	var bn string
	if err = controller.DecryptRequest[string](ctx, &bn); err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数解析错误"))
		return
	}
	ser := BookService{}
	if err = ser.UpdateBookName(id, bn); err != nil {
		controller.ForbiddenResp(ctx, errors.New("更新失败"))
		return
	}
	controller.EncryptAnyResp(ctx, "ok")
}

// CreateAttachment 插入附件
func (b *BookController) CreateAttachment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	uid := ctx.MustGet("uid").(uint)
	//TODO 检查书籍是否存在
	file, err := ctx.FormFile("file")
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("获取文件失败"))
		return
	}
	// 检查文件类型
	if utils.IsForbiddenFile(file.Filename) {
		controller.ForbiddenResp(ctx, errors.New("文件禁止上传"))
		return
	}
	fileStore := models.NewFileMode(file.Filename, file.Size, uid)
	// 保存文件
	if err = ctx.SaveUploadedFile(file, path.Join(config.Cfg.UploadDir, fileStore.Filename)); err != nil {
		controller.ForbiddenResp(ctx, errors.New("获取文件失败"))
		return
	}
	// 保存信息
	ser := BookService{}
	if err = ser.InsertAttachment(models.BookAttachment{BookID: uint(id), FileStoreModel: *fileStore}); err != nil {
		controller.ForbiddenResp(ctx, errors.New("保存失败"))
		return
	}
	controller.SuccessResp(ctx, fileStore.Filename)
}

// DelAttachment 删除附件
func (b *BookController) DelAttachment(ctx *gin.Context) {
	ser := BookService{}
	if err := ser.DelAttachment(ctx.Param("filename")); err != nil {
		controller.ForbiddenResp(ctx, errors.New("删除失败"))
		return
	}
	controller.SuccessResp(ctx, "ok")
}

// DownloadAttachment 下载附件
func (b *BookController) DownloadAttachment(ctx *gin.Context) {
	fn := ctx.Param("filename")
	if fn == "" {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	ser := BookService{}
	ft, err := ser.GetAttachment(fn)
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("文件不存在或已删除"))
		return
	}
	ctx.FileAttachment(path.Join(config.Cfg.UploadDir, fn), ft.OriginName)
}

// UpdateNote 更新note
func (b *BookController) UpdateNote(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	var note string
	if err = controller.DecryptRequest[string](ctx, &note); err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数解析错误"))
		return
	}
	ser := BookService{}
	if err = ser.UpdateNote(id, note); err != nil {
		controller.ForbiddenResp(ctx, errors.New("更新失败"))
		return
	}
	controller.EncryptAnyResp(ctx, "ok")
}

// UpdateSourceLink 更新SourceLink
func (b *BookController) UpdateSourceLink(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数错误"))
		return
	}
	var note string
	if err = controller.DecryptRequest[string](ctx, &note); err != nil {
		controller.ForbiddenResp(ctx, errors.New("参数解析错误"))
		return
	}
	ser := BookService{}
	if err = ser.UpdateSourceLink(id, note); err != nil {
		controller.ForbiddenResp(ctx, errors.New("更新失败"))
		return
	}
	controller.EncryptAnyResp(ctx, "ok")
}
