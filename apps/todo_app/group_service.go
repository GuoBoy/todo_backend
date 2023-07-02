package todo_app

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"todo_backend/db"
	"todo_backend/models"
	"todo_backend/service"
)

// GroupService 分组服务
type GroupService struct {
	D *models.GroupModel
}

// Create 新增分组
func (itm *GroupService) Create() (int64, error) {
	dt := time.Now().Unix()
	return dt, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&itm.D).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, itm.D.UserID, dt); err != nil {
			return err
		}
		return nil
	})
}

// Delete 删除分组
func (itm *GroupService) Delete(gid string, uid uint) (int64, error) {
	dt := time.Now().Unix()
	return dt, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.GroupModel{}, "id=?", gid).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, uid, dt); err != nil {
			return err
		}
		return nil
	})
}

// ForceDelete 彻底删除
func (itm *GroupService) ForceDelete(gid string, uid uint) (int64, error) {
	dt := time.Now().Unix()
	return dt, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&models.GroupModel{}, "id=?", gid).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Delete(&models.ItemModel{}, "group_id=?", gid).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("错误", err)
			return err
		}
		if err := service.UpdateUserTimeByID(tx, uid, dt); err != nil {
			return err
		}
		return nil
	})
}

// Update 更新分组名称
func (itm *GroupService) Update(uid uint, gid, name string) (int64, error) {
	dt := time.Now().Unix()
	return dt, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&models.GroupModel{}, "id=?", gid).Update("name", name).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, uid, dt); err != nil {
			return err
		}
		return nil
	})
}

// All 获取所有分组
func (itm *GroupService) All(uid uint) ([]*models.GroupRespModel, error) {
	var (
		resp   []*models.GroupRespModel
		result []*models.GroupModel
	)
	if err := db.DB.Find(&result, "user_id=?", uid).Error; err != nil {
		return nil, err
	}
	for _, group := range result {
		resp = append(resp, &models.GroupRespModel{
			ID:   group.ID,
			Name: group.Name,
		})
	}
	return resp, nil
}

// CheckGroupByID 检查分组是否存在
func CheckGroupByID(gid string) error {
	if gid == "" {
		return errors.New("小组不存在")
	}
	return db.DB.First(&models.GroupModel{}, "id=?", gid).Error
}
