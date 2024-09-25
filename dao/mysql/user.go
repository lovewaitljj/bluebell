package mysql

import (
	"bluebell/models"
	"errors"
)

// CheckUserExist 判断用户是否存在(用户名)
func CheckUserExist(username string) error {
	var count int
	if err := db.Table(models.User{}.TableName()).Select("COUNT(*)").Where("username = ?", username).Find(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
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
