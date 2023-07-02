package models

import (
	"gorm.io/gorm"
)

// User 存储用户模型
type User struct {
	*gorm.Model
	Username   string `json:"username" gorm:"username;unique" binding:"required"`
	Password   string `json:"password" gorm:"password" binding:"required"`
	UpdateTime int64  `json:"update_time" gorm:"update_time"`
	Enabled    bool   `json:"enabled" gorm:"enabled;default:true"`
}

// ResponseUser 用户列表响应用户模型
type ResponseUser struct {
	ID         uint   `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Enabled    bool   `json:"enabled"`
}

// UserForm 登录注册用户表单
type UserForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckUpdateForm struct {
	Userid         uint  `json:"userid" binding:"required"`
	LastUpdateTime int64 `json:"last_update_time" binding:"required"`
}

// PwdForm 密码更新表单
type PwdForm struct {
	O string `json:"o,omitempty" binding:"required"`
	N string `json:"n,omitempty" binding:"required"`
}
