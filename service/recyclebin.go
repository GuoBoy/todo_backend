package service

import (
	"gorm.io/gorm"
	"time"
	"todo_backend/config"
	"todo_backend/db"
	"todo_backend/models"
)

type RecycleGroupItem struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	DeletedAt string `json:"deleted_at"`
}

type RecycleTodoItem struct {
	ID        uint   `json:"id"`
	Detail    string `json:"detail" gorm:"not null"`
	DeletedAt string `json:"deleted_at"`
}

type RecycleService struct {
	Groups []*RecycleGroupItem `json:"groups,omitempty"`
	Items  []*RecycleTodoItem  `json:"items,omitempty"`
}

func (r *RecycleService) All(uid uint) error {
	var (
		groups []*models.GroupModel
		items  []*models.ItemModel
	)
	if err := db.DB.Raw("SELECT * FROM group_models WHERE deleted_at IS NOT NULL AND user_id=?", uid).Scan(&groups).Error; err != nil {
		return err
	}
	for _, group := range groups {
		r.Groups = append(r.Groups, &RecycleGroupItem{
			ID:        group.ID,
			Name:      group.Name,
			DeletedAt: group.DeletedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	if err := db.DB.Raw("SELECT * FROM item_models WHERE deleted_at IS NOT NULL AND user_id=?", uid).Scan(&items).Error; err != nil {
		return err
	}
	for _, item := range items {
		r.Items = append(r.Items, &RecycleTodoItem{
			ID:        item.ID,
			Detail:    item.Detail,
			DeletedAt: item.DeletedAt.Time.Format(config.TimeLayout),
		})
	}
	return nil
}

func (r *RecycleService) GroupRecoverByID(gid string, uid uint) (int64, error) {
	dt := time.Now().Unix()
	return dt, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE group_models SET deleted_at=NULL WHERE id=?", gid).Error; err != nil {
			return err
		}
		if err := UpdateUserTimeByID(tx, uid, dt); err != nil {
			return err
		}
		return nil
	})
}

func (r *RecycleService) ItemRecoverByID(tid string, uid uint) (int64, error) {
	dt := time.Now().Unix()
	return dt, db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE item_models SET deleted_at=NULL WHERE id=?", tid).Error; err != nil {
			return err
		}
		if err := UpdateUserTimeByID(tx, uid, dt); err != nil {
			return err
		}
		return nil
	})
}
