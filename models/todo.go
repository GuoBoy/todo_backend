package models

import (
	"gorm.io/gorm"
)

// GroupModel 分组模型
type GroupModel struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"user_id"`
	Name   string `json:"name" gorm:"name" binding:"required"`
}

// GroupRespModel 响应的分组模型
type GroupRespModel struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ItemModel todo模型
type ItemModel struct {
	*gorm.Model
	Detail   string `json:"detail" gorm:"not null"`
	Checked  bool   `json:"checked" gorm:"checked"` // item 是否checked状态
	Hashcode string `json:"hashcode" gorm:"hashcode"`
	GroupId  uint   `json:"group_id" gorm:"group_id"`
	UserID   uint   `json:"user_id" gorm:"user_id"`
}

// ItemResponseModel item响应模型
type ItemResponseModel struct {
	ID        uint   `json:"id,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at"`
	Detail    string `json:"detail,omitempty"`
	Checked   bool   `json:"checked"`
	Hashcode  string `json:"hashcode" gorm:"hashcode"`
	GroupId   uint   `json:"group_id"`
	UserID    uint   `json:"user_id"`
}

// ItemForm todo新增表单模型 / 更新模型，更新时需要hashcode
type ItemForm struct {
	GroupID  uint   `json:"group_id" binding:"required"`
	Detail   string `json:"detail" binding:"required"`
	Hashcode string `json:"hashcode"`
}

// DoItemDelForm 删除item模型
type DoItemDelForm struct {
	Uid uint `json:"uid,omitempty" binding:"required"`
	Tid uint `json:"tid,omitempty" binding:"required"`
}
