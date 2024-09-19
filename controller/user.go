package controller

import (
	"bluebell/models"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	//1.获取参数
	var req models.RegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		zap.L().Error("register with invalid param", zap.Error(err))
		//判断err是否为validator类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//2.业务逻辑
	err = service.Register(&req)
	if err != nil {
		zap.L().Error("register with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	//3.返回相应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
