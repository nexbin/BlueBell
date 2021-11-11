package middleware

import (
	"BlueBell/controllers"
	"BlueBell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthorizationMiddleware 基于JWT的认证中间件
func JWTAuthorizationMiddleware() func(c *gin.Context) {
	// 客户端三种方式携带token: 请求头、请求体、URI
	// 这里是放在请求头中，Bearer开头
	// Authorization:Bearer xxx.xxx.xxx
	// 具体业务实现
	return func(c *gin.Context) {
		authHead := c.Request.Header.Get("Authorization")
		if authHead == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHead, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 解析Token
		token, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userId信息保存在请求上下文中
		c.Set(controllers.ContextUserIdKey, token.UserId)
		c.Next() // 后续的处理函数可以通过C.Get(ContextUserIdKey)来获取userid
	}
}
