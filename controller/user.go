package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// RegisterHandler 注册
func RegisterHandler(c *gin.Context) {
	//1.获取参数
	var req models.RegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error("register with invalid param", zap.Error(err))
		//判断err是否为validator类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务逻辑
	err = service.Register(&req)
	if err != nil {
		zap.L().Error(" service.Register failed", zap.Error(err))
		if errors.Is(err, mysql.ErrUserExists) {
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回相应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	//1.获取参数
	var req models.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		//判断err是否为validator类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务逻辑
	if err = service.Login(&req); err != nil {
		zap.L().Error("login failed", zap.String("username", req.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrUserNotExists) {
			ResponseError(c, CodeUserNotExist)
		} else if errors.Is(err, mysql.ErrInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回相应
	ResponseSuccess(c, nil)
}
