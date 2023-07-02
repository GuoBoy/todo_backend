package utils

import (
	"path"
	"time"
	"todo_backend/config"
	"todo_backend/consts"
)

// IsForbiddenFile 验证文件是否禁止上传
func IsForbiddenFile(filename string) bool {
	ext := path.Ext(filename)
	for _, e := range config.Cfg.ExtBlacklist {
		if e == ext {
			return true
		}
	}
	return false
}

// FormatTime2Standard 格式化时间
func FormatTime2Standard(ti time.Time) string {
	return ti.Format(consts.StandardTimeLayout)
}
