package todo_app

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"todo_backend/db"
	"todo_backend/models"
	"todo_backend/service"
	"todo_backend/utils"
)

type ItemService struct {
	UserID uint
	D      *models.ItemModel
}

// Create 创建todo
// return 用户更新时间、err
func (t *ItemService) Create() (int64, error) {
	ut := time.Now().Unix()
	return ut, db.DB.Transaction(func(tx *gorm.DB) error {
		t.D.Hashcode = utils.MakeMd5([]byte(t.D.Detail))
		if err := tx.Create(&t.D).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, t.UserID, ut); err != nil {
			return err
		}
		return nil
	})
}

func (t *ItemService) Delete(tid string, uid uint) (int64, error) {
	ut := time.Now().Unix()
	return ut, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.ItemModel{}, tid).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, uid, ut); err != nil {
			return err
		}
		return nil
	})
}

func (t *ItemService) ForceDelete(tid string, uid uint) (int64, error) {
	ut := time.Now().Unix()
	return ut, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&models.ItemModel{}, tid).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, uid, ut); err != nil {
			return err
		}
		return nil
	})
}

// GetAllByUid 获取用户所有待办
func (*ItemService) GetAllByUid(uid uint) ([]*models.ItemResponseModel, error) {
	var (
		items     []*models.ItemModel
		respItems []*models.ItemResponseModel
	)
	if err := db.DB.Find(&items, "user_id=?", uid).Order("updated_at DESC").Error; err != nil {
		return respItems, err
	}
	for _, item := range items {
		respItems = append(respItems, &models.ItemResponseModel{
			ID:        item.ID,
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
			Detail:    item.Detail,
			Hashcode:  item.Hashcode,
		})
	}
	return respItems, nil
}

func (*ItemService) GetAllByGroupID(gid string) ([]*models.ItemResponseModel, error) {
	var (
		items     []*models.ItemModel
		respItems []*models.ItemResponseModel
	)
	if err := db.DB.Find(&items, "group_id=?", gid).Order("updated_at DESC").Error; err != nil {
		return nil, err
	}
	for _, item := range items {
		respItems = append(respItems, &models.ItemResponseModel{
			ID:        item.ID,
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
			Detail:    item.Detail,
			Checked:   item.Checked,
			Hashcode:  item.Hashcode,
			GroupId:   item.GroupId,
		})
	}
	return respItems, nil
}

// Update 更新item服务
func (t *ItemService) Update(uid uint, tid, hashcode, content string) (int64, error) {
	ut := time.Now().Unix()
	return ut, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&t.D, map[string]string{"id": tid, "hashcode": hashcode}).Updates(map[string]interface{}{"detail": content, "hashcode": utils.MakeMd5([]byte(content))}).Error; err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("未报错")
		if err := tx.Model(&models.User{}).Where("id=?", uid).Update("update_time", ut).Error; err != nil {
			return err
		}
		return nil
	})
}

func (t *ItemService) UpdateChecked(uid, tid int, checked bool) (int64, error) {
	ut := time.Now().Unix()
	return ut, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&t.D, tid).Update("checked", checked).Error; err != nil {
			return err
		}
		if err := service.UpdateUserTimeByID(tx, uint(uid), ut); err != nil {
			return err
		}
		return nil
	})
}
