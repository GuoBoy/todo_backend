package models

// AdminUserModel 管理员用户登录表单
type AdminUserModel struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// AdminUserEnabledForm 修改用户可用性表单
type AdminUserEnabledForm struct {
	Uid     int  `json:"uid" binding:"required"`
	Enabled bool `json:"enabled" binding:"required"`
}

// ChangeItemDetailForm 修改detail
type ChangeItemDetailForm struct {
	ID     int    `json:"id" binding:"required"`
	Detail string `json:"detail" binding:"required"`
}

// ChangeCheckedForm 切换check状态
type ChangeCheckedForm struct {
	Uid     int  `json:"uid,omitempty" binding:"required"`
	ID      int  `json:"ID,omitempty" binding:"required"`
	Checked bool `json:"checked,omitempty" binding:"required"`
}
