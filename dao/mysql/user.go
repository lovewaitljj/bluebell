package mysql

import (
	"bluebell/models"
	"bluebell/pkg/md5"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrUserExists      = errors.New("用户已存在")
	ErrUserNotExists   = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 判断用户是否存在(用户名)
func CheckUserExist(username string) error {
	var count int
	if err := db.Table(models.User{}.TableName()).Select("COUNT(*)").Where("username = ?", username).Find(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExists
	}
	return nil
}

// CreateUser 插入用户
func CreateUser(user *models.User) error {
	if err := db.Create(user).Debug().Error; err != nil {
		return err
	}
	return nil
}

// Login 登录
func Login(p *models.User) error {
	//查询是否存在用户
	var user models.User
	if err := db.Select("username,password").Where("username = ?", p.Username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotExists
		}
		return err
	}
	//判断密码是否正确
	if md5.Encrypt(p.Password) != user.Password {
		return ErrInvalidPassword
	}
	return nil
}
