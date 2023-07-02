package booknote_app

import (
	"github.com/gin-gonic/gin"
	"todo_backend/midddleware"
)

func BookNoteRouter(r *gin.Engine) {
	bnc := BookController{}
	book := r.Group("/book", midddleware.AccessLog())
	{
		// 免用户验证
		book.GET("/attachment/:filename", bnc.DownloadAttachment) // 附件下载
		// 需权限验证
		book.Use(midddleware.RequestAuth())
		// book
		book.GET("", bnc.GetAll)
		book.POST("", bnc.Create)
		book.DELETE("/:id", bnc.Del)
		book.POST("/:id", bnc.UpdateBookName)
		// note
		book.POST("/note/:id", bnc.UpdateNote) //书籍id更新note
		book.POST("/sourcelink/:id", bnc.UpdateSourceLink)
		// attachment
		book.POST("/attachment/:id", bnc.CreateAttachment)
		book.DELETE("attachment/:filename", bnc.DelAttachment)
	}
}
