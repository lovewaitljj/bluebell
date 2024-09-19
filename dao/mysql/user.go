package mysql

import (
	"bluebell/models"
	"errors"
	"gorm.io/gorm"
)

// CheckUserExist 判断用户是否存在(用户名)
func CheckUserExist(username string) (bool, error) {
	user := new(models.User)
	if err := db.Where("username = ?", username).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateUser 插入用户
func CreateUser(user *models.User) error {
	if err := db.Create(user).Debug().Error; err != nil {
		return err
	}
	return nil
}
