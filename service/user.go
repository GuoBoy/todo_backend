package service

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"todo_backend/config"
	"todo_backend/db"
	"todo_backend/models"
	"todo_backend/utils"
)

type User struct {
	D *models.User
}

// Create 创建用户
// param 用户名、密码
// return err
func (u *User) Create() error {
	if err := db.DB.First(&u.D, "username=?", u.D.Username).Error; err == nil {
		return errors.New("用户已存在")
	}
	u.D.Password = utils.StrMd5(u.D.Password)
	u.D.UpdateTime = time.Now().Unix()
	return db.DB.Create(&u.D).Error
}

// VerifyByUaPW 用户验证
// param 用户名、密码
// return err
func (u *User) VerifyByUaPW() error {
	u.D.Password = utils.StrMd5(u.D.Password)
	return db.DB.Take(&u.D, map[string]string{"username": u.D.Username, "password": u.D.Password}).Error
}

// CheckUpdateTime 验证更新时间，
// param 用户id, 更新时间
// return 是否需要更新, err
func (u *User) CheckUpdateTime(uid uint, ut int64) (bool, error) {
	if err := db.DB.First(&u.D, uid).Error; err != nil {
		return false, err
	}
	if u.D.UpdateTime > ut {
		return true, nil
	}
	return false, nil
}

func GetUserUpdateTime(uid uint) (int64, error) {
	var u models.User
	if err := db.DB.First(&u, uid).Error; err != nil {
		return 0, err
	}
	return u.UpdateTime, nil
}

// UpdatePwd 修改密码
func (u *User) UpdatePwd(uid uint, pwd models.PwdForm) error {
	tx := db.DB.First(&u.D, "id=? AND password=?", uid, utils.MakeMd5([]byte(pwd.O)))
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Update("password", utils.MakeMd5([]byte(pwd.N))).Error
}

// UpdatePwdToDefault 修改密码
func (u *User) UpdatePwdToDefault(uid int) error {
	return db.DB.First(&u.D, uid).Update("password", utils.StrMd5(config.Cfg.DefaultResetPwd)).Error
}

// DeleteUserByID 删除用户
func (u *User) DeleteUserByID(id int) error {
	return db.DB.Unscoped().Delete(&u.D, id).Error
}

// SetUserEnabled 设置用户可用状态
func (u *User) SetUserEnabled(form models.AdminUserEnabledForm) error {
	return db.DB.First(&models.User{}, form.Uid).Update("enabled", form.Enabled).Error
}

// All 用户列表
func (*User) All() ([]*models.ResponseUser, error) {
	var (
		users []*models.User
		resU  []*models.ResponseUser
	)
	if err := db.DB.Find(&users).Error; err != nil {
		return resU, err
	}
	for _, user := range users {
		resU = append(resU, &models.ResponseUser{
			ID:         user.ID,
			Username:   user.Username,
			CreatedAt:  user.CreatedAt.Format("2006-01-02 13:04:05"),
			UpdatedAt:  user.UpdatedAt.Format("2006-01-02 13:04:05"),
			UpdateTime: time.Unix(user.UpdateTime, 0).Format("2006-01-02 13:04:05"),
			Enabled:    user.Enabled,
		})
	}
	return resU, nil
}

// UpdateUserTimeByID 更新用户时间
func UpdateUserTimeByID(tx *gorm.DB, uid uint, ndt int64) error {
	return tx.First(&models.User{}, "id = ?", uid).Update("update_time", ndt).Error
}

// FeedBackService 反馈信息服务
type FeedBackService struct {
}

// StoreFeedBack 保存反馈信息
func (f FeedBackService) StoreFeedBack(form *models.FeedBackModel) error {
	return db.DB.Create(&form).Error
}
