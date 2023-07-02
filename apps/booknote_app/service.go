package booknote_app

import (
	"todo_backend/config"
	"todo_backend/db"
	"todo_backend/models"
)

type BookService struct {
	D *models.BookNote
}

func (b *BookService) Add() error {
	return db.DB.Create(&b.D).Error
}

// GetAll 获取所有
func (b *BookService) GetAll() ([]*models.BookNoteListItem, int, error) {
	var (
		res          []*models.BookNoteListItem
		tempBookNote []*models.BookNote
		tempAttachs  []*models.BookAttachment
	)
	if err := db.DB.Find(&tempBookNote).Error; err != nil {
		return nil, 0, err
	}
	for _, note := range tempBookNote {
		if err := db.DB.Find(&tempAttachs, "book_id=?", note.ID).Error; err != nil {
			return nil, 0, err
		}
		res = append(res, &models.BookNoteListItem{
			ID:          note.ID,
			BookName:    note.BookName,
			Note:        note.Note,
			SourceLink:  note.SourceLink,
			CreatedAt:   note.CreatedAt.Format(config.TimeLayout),
			UpdatedAt:   note.CreatedAt.Format(config.TimeLayout),
			Attachments: tempAttachs,
		})
	}
	//tx := db.DB.Raw("SELECT * FROM book_notes n LEFT JOIN book_attachments  a ON n.id=a.book_id").Scan(&res)
	return res, len(res), nil
}

// GetByID 获取一个
func (b *BookService) GetByID(id int) (models.BookNoteListItem, int64, error) {
	var res models.BookNoteListItem
	tx := db.DB.Raw("SELECT * FROM book_notes n LEFT JOIN book_attachments  a ON n.id=a.book_id WHERE n.id=?", id).Scan(&res)
	return res, tx.RowsAffected, tx.Error
}

// DelByID 删除：非真实删除
func (b *BookService) DelByID(id int) error {
	return db.DB.Delete(&models.BookNote{}, id).Error
}

// UpdateBookName 更新书籍名称
func (b *BookService) UpdateBookName(bid int, link string) error {
	return db.DB.Take(&models.BookNote{}, bid).Update("book_name", link).Error
}

// InsertAttachment 插入附件
func (b *BookService) InsertAttachment(att models.BookAttachment) error {
	return db.DB.Create(att).Error
}

// DelAttachment 删除附件
func (b *BookService) DelAttachment(filename string) error {
	return db.DB.Delete(&models.BookAttachment{}, "filename=?", filename).Error
}

func (*BookService) GetAttachment(filename string) (*models.BookAttachment, error) {
	var res models.BookAttachment
	err := db.DB.Take(&res, "filename=?", filename).Error
	return &res, err
}

func (b *BookService) UpdateNote(bid int, content string) error {
	return db.DB.Take(&models.BookNote{}, bid).Update("note", content).Error
}

func (b *BookService) UpdateSourceLink(bid int, link string) error {
	return db.DB.Take(&models.BookNote{}, bid).Update("source_link", link).Error
}
