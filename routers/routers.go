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

	v1 := r.Group("/api/v1")
	{
		// 用户登录
		v1.POST("/login", controllers.LoginHandler)
		// 处理注册
		v1.POST("/signup", controllers.SignUpHandler)

		v1.Use(middleware.JWTAuthorizationMiddleware())
		{
			// 获取社区列表
			v1.GET("/community", controllers.CommunityHandler)
			// 根据id获取社区具体信息
			v1.GET("/community/:id", controllers.CommunityDetailHandler)

			// 发布帖子
			v1.POST("/post", controllers.CreatePostHandler)
			// 帖子详情，点击帖子，进入
			v1.GET("/post/:id", controllers.GetPostDetailHandler)
			// 获取帖子列表
			v1.GET("/posts", controllers.GetPostList)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
