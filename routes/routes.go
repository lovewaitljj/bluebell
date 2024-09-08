package routes

import (
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
	return r
}
