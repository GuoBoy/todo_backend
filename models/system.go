package models

import "gorm.io/gorm"

// AppVersion 各端应用版本
type AppVersion struct {
	gorm.Model
	PlatformName string              `json:"platform_name"` // 平台
	Histories    []AppHistoryVersion `gorm:"foreignKey:AppID" json:"histories"`
}

// AppHistoryVersion 平台更新历史
type AppHistoryVersion struct {
	gorm.Model
	VersionName  string `json:"version_name,omitempty"` // 版本号
	VersionValue int64  `json:"version_value,omitempty"`
	Description  string `json:"description,omitempty"` // 更新说明
	AppID        uint   `json:"app_id,omitempty"`
}

// PaginationQuery 分页查询参数
type PaginationQuery struct {
	Size int `form:"size"`
	Page int `form:"page"`
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total int64       `json:"total"` // 数据量
	Data  interface{} `json:"data"`  // 数据
}

// UserTokenInfo token保存的用户信息
type UserTokenInfo struct {
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
}
