package service

import (
	"todo_backend/db"
	"todo_backend/models"
)

type AppService struct {
}

// CreateApp 创建app
func (AppService) CreateApp(app models.AppVersion) (uint, error) {
	if err := db.DB.Create(&app).Error; err != nil {
		return 0, err
	}
	return app.ID, nil
}

// GetAllApp 获取所有app
func (AppService) GetAllApp() ([]*models.AppVersion, error) {
	var res []*models.AppVersion
	//if err := db.DB.Model(&models.AppVersion{}).Preload("Histories").Find(&res).Error; err != nil {
	if err := db.DB.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateAppHistory 添加app history
func (AppService) CreateAppHistory(app models.AppHistoryVersion) (uint, error) {
	if err := db.DB.Create(&app).Error; err != nil {
		return 0, err
	}
	return app.ID, nil
}

// GetAppHistoryByID 获取app更新历史
func (AppService) GetAppHistoryByID(id string) ([]*models.AppHistoryVersion, error) {
	var res []*models.AppHistoryVersion
	if err := db.DB.Find(&res, "app_id=?", id).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetVersion 获取版本信息
func GetVersion(plat string) (int64, string, error) {
	var his models.AppHistoryVersion
	tx := db.DB.Raw("SELECT * from app_history_versions as a LEFT JOIN app_versions as b on a.app_id=b.id WHERE b.platform_name = ? ORDER BY a.version_value DESC LIMIT 1", plat).Scan(&his)
	if tx.Error != nil {
		return 0, "", tx.Error
	}
	return his.VersionValue, his.VersionName, nil
}
