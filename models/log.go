package models

import "time"

// AccessLogModel 访问日志模型
type AccessLogModel struct {
	ID         uint          `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt  time.Time     `json:"created_at"`           // 访问时间
	ClientIP   string        `json:"client_ip,omitempty"`  // 客户端IP
	UserAgent  string        `json:"user_agent,omitempty"` // ua
	Method     string        `json:"method"`
	Path       string        `json:"path,omitempty"`        // 访问地址
	StatusCode int           `json:"status_code,omitempty"` // 响应状态码
	UID        uint          `json:"uid" gorm:""`           // 访问用户
	Username   string        `json:"username"`              // 用户名
	Latency    time.Duration `json:"latency"`               // 处理时长
}
