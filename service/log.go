package service

import (
	"todo_backend/db"
	"todo_backend/models"
)

// LogService 日志服务
type LogService struct {
}

// AddLog 添加日志
func AddLog(lg models.AccessLogModel) error {
	return db.DB.Create(&lg).Error
}

// GetAllLogs 获取日志
func (s LogService) GetAllLogs(pa models.PaginationQuery) (models.PaginationResult, error) {
	var (
		pag models.PaginationResult
		res []*models.AccessLogModel
	)
	tx := db.DB.Model(models.AccessLogModel{}).Order("created_at DESC")
	if tx.Error != nil {
		return pag, tx.Error
	}
	// 总数
	tx.Count(&pag.Total)
	if err := tx.Limit(pa.Size).Offset((pa.Page - 1) * pa.Size).Find(&res).Error; err != nil {
		return pag, err
	}
	pag.Data = res
	return pag, nil
}
