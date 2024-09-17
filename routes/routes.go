package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "okk")
	})
	//业务逻辑路由
	r.POST("/register", controller.RegisterHandler)

	return r
}
