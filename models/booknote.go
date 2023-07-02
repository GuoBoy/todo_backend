package models

import (
	"gorm.io/gorm"
)

// BookNote 书籍日志模型
type BookNote struct {
	gorm.Model
	BookName   string `json:"book_name,omitempty"  binding:"required"` //书籍名称
	Note       string `json:"note"`                                    // 笔记
	SourceLink string `json:"source_link"`                             // 来源网址
}

// BookAttachment 书籍附件,一个附件对应一本书籍，一个书籍可多个附件
type BookAttachment struct {
	BookID uint `json:"book_id"`
	FileStoreModel
}

// BookNoteListItem 返回书籍日志模型
type BookNoteListItem struct {
	ID          uint              `json:"id"`
	BookName    string            `json:"book_name,omitempty"` //书籍名称
	Note        string            `json:"note"`                // 笔记
	SourceLink  string            `json:"source_link"`         // 来源网址
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	Attachments []*BookAttachment `json:"attachments"`
}
