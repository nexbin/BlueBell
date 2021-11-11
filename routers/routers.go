package routers

import (
	"BlueBell/controllers"
	"BlueBell/logger"
	"BlueBell/middleware"
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
	r.POST("/signup", controllers.SignUpHandler)

	// 用户登录
	r.POST("/login", controllers.LoginHandler)

	// 测试Token
	r.GET("/ping", middleware.JWTAuthorizationMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户,判断当前用户是否是登录的用户（请求头中是否有有效的jwtToken）
		_, exists := c.Get("userId")
		if exists {
			c.String(http.StatusOK, "hello,token carrays.\n")
		}
		// 否则
	})

	return r
}
