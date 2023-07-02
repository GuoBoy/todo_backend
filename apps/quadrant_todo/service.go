package quadrant_todo

import (
	"strconv"
	"todo_backend/db"
	"todo_backend/models"
	"todo_backend/utils"
)

// Service 象限todo model
type Service struct {
}

// Create 创建 q todoItem
// return 用户更新时间、err
func (s *Service) Create(form *models.QTodoModel) error {
	form.Hashcode = utils.StrMd5(form.Detail)
	return db.DB.Create(&form).Error
}

// Delete 删除q todo服务
func (s *Service) Delete(tid string, uid uint) error {
	item := models.QTodoModel{UserID: uid}
	tidr, _ := strconv.ParseUint(tid, 10, 32)
	item.ID = uint(tidr)
	return db.DB.Delete(&item).Error
}

// GetAllByQuadrantType 获取用户象限的待办
func (*Service) GetAllByQuadrantType(uid, quadrant uint) (items []*models.ResponseQTodoModel, err error) {
	var temp []*models.QTodoModel
	if err = db.DB.Where(map[string]any{"user_id": uid, "quadrant_type": quadrant}).Order("updated_at DESC").Find(&temp).Error; err != nil {
		return
	}

	for _, item := range temp {
		items = append(items, &models.ResponseQTodoModel{
			ID:           item.ID,
			UpdatedAt:    utils.FormatTime2Standard(item.UpdatedAt),
			Detail:       item.Detail,
			QuadrantType: item.QuadrantType,
			Hashcode:     item.Hashcode,
		})
	}
	return
}

// UpdateDetail 更新详情
func (s *Service) UpdateDetail(id, detail string) (item *models.QTodoModel, err error) {
	err = db.DB.First(&item, id).Updates(map[string]interface{}{"detail": detail, "hashcode": utils.MakeMd5([]byte(detail))}).Error
	return
}
