package models

import "gorm.io/gorm"

// 象限值
const (
	FirstQuadrant = iota + 1
	SecondQuadrant
	ThirdQuadrant
	FourthQuadrant
)

// QTodoModel 象限todo
type QTodoModel struct {
	gorm.Model
	Detail       string `json:"detail" binding:"required"`        // 内容
	QuadrantType uint   `json:"quadrant_type" binding:"required"` // 类型
	UserID       uint   `json:"user_id" gorm:"user_id"`
	Hashcode     string `json:"hashcode" gorm:"hashcode"`
}

// ResponseQTodoModel 象限todo
type ResponseQTodoModel struct {
	ID           uint   `json:"id"`
	UpdatedAt    string `json:"updated_at"`
	Detail       string `json:"detail" binding:"required"`        // 内容
	QuadrantType uint   `json:"quadrant_type" binding:"required"` // 类型
	UserID       uint   `json:"user_id" gorm:"user_id"`
	Hashcode     string `json:"hashcode" gorm:"hashcode"`
}
