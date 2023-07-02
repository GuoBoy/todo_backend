package admin

import (
	"todo_backend/db"
	"todo_backend/models"
	"todo_backend/utils"
)

// AllItems 获取所有item
func AllItems() (res []*models.ItemResponseModel, err error) {
	err = db.DB.Model(&models.ItemModel{}).Find(&res).Error
	return res, err
}

func UpdateItemDetail(id int, detail string) error {
	return db.DB.Model(&models.ItemModel{}).Where("id = ?", id).Updates(map[string]interface{}{"detail": detail, "hashcode": utils.MakeMd5([]byte(detail))}).Error
}

func DelItem(id int) error {
	return db.DB.Delete(&models.ItemModel{}, id).Error
}
