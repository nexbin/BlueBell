package routers

import (
	"BlueBell/controllers"
	"BlueBell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})

	// 处理注册
	r.GET("/signup", controllers.SignUpHandler)

	return r
}
