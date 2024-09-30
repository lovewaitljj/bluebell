package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/md5"
	sf "bluebell/pkg/snowflake"
)

func Register(req *models.RegisterReq) error {
	//1.判断用户存不存在、
	err := mysql.CheckUserExist(req.Username)
	if err != nil {
		return err
	}
	//2.生成uid并插入db
	userId := sf.GenID()
	user := &models.User{
		UserID:   uint64(userId),
		Username: req.Username,
		Password: md5.Encrypt(req.Password),
		Email:    req.Email,
		Gender:   req.Gender,
	}
	err = mysql.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(req *models.LoginReq) error {
	//1.用户名是否存在
	//2.密码是否正确
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}
	return mysql.Login(user)
}
