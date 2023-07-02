package service

import (
	"todo_backend/db"
	"todo_backend/models"
)

// FileService 文件服务
type FileService struct {
}

// Create 保存文件
func (f *FileService) Create(file *models.FileStoreModel) error {
	return db.DB.Create(&file).Error
}

// GetInfo 获取文件信息
func (f *FileService) GetInfo(filename string) (models.FileStoreModel, error) {
	var fm models.FileStoreModel
	if err := db.DB.First(&fm, "filename=?", filename).Error; err != nil {
		return models.FileStoreModel{}, err
	}
	return fm, nil
}

// Del 删除文件
func (f *FileService) Del(filename string) error {
	if err := db.DB.Delete(&models.FileStoreModel{}, "filename=?", filename).Error; err != nil {
		return err
	}
	return nil
}
