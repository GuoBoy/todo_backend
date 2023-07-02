package models

import "gorm.io/gorm"

// FeedBackModel 反馈模型
type FeedBackModel struct {
	gorm.Model
	Message  string `json:"message" binding:"required"`
	Uid      uint   `json:"uid"`
	Images   string `json:"images"`                        // imgID1|imgID2
	Useful   bool   `json:"useful" gorm:"default:true"`    // 反馈是否有价值
	Finished bool   `json:"finished" gorm:"default:false"` // 反馈是否结束
}
