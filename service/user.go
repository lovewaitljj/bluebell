package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/md5"
	sf "bluebell/pkg/snowflake"
	"errors"
)

func Register(req *models.RegisterReq) error {
	//1.判断用户存不存在、
	exist, err := mysql.CheckUserExist(req.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户已存在")
	}
	//2.生成uid并插入db
	userId := sf.GenID()
	user := &models.User{
		UserID:   uint64(userId),
		Username: req.Username,
		Password: md5.MD5Encrypt((req.Password)),
	}
	err = mysql.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
