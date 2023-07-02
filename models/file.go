package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// FileStoreModel 存储文件模型
type FileStoreModel struct {
	Filename   string    `json:"filename,omitempty" gorm:"primaryKey"` // 文件名称，md5(原文件名称+时间戳)
	OriginName string    `json:"origin_name"`
	Filesize   int64     `json:"filesize"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"` // 创建时间
	Uid        uint      `json:"uid"`                                        // 用户id
}

func makeMd5(st string) string {
	temp := md5.Sum([]byte(st))
	return hex.EncodeToString(temp[:])
}

func NewFileMode(originName string, size int64, uid uint) *FileStoreModel {
	tNow := time.Now()
	return &FileStoreModel{
		Filename:   makeMd5(fmt.Sprintf("%s+%d", originName, tNow.UnixMicro())),
		OriginName: originName,
		Filesize:   size,
		CreatedAt:  tNow,
		Uid:        uid,
	}
}
