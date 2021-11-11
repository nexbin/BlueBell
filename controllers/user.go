package controllers

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Sign up with invalid param,", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		// 如果错误类型不为validator.ValidationErrors类型，说明可能是其他错误,比如序列化错误等
		if !ok {
			c.JSON(http.StatusOK, &gin.H{
				"msg": err.Error(),
			})
		} else {
			// 否则为validator.ValidationErrors，返回json格式的新error
			c.JSON(http.StatusOK, &gin.H{
				"msg": errs.Translate(trans),
			})
		}
		return
	}
	// 手动对请求参数进行详细的校验
	{
		//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		//	zap.L().Error("Sign up with invalid param.")
		//	c.JSON(http.StatusOK, &gin.H{
		//		"msg": "请求参数有误",
		//	})
		//	return
		//}
	}

	// 2.业务处理
	if err := logic.UserSignUp(p); err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, &gin.H{"Status": "ok"})
}

// LoginHandler 处理用户登录请求的函数
func LoginHandler(c *gin.Context) {
	// 1.获得请求参数（用户账号和密码）,并进行参数校验
	u := new(models.ParamLogin)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 如果是validator.ValidationErrors类型
		if err, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{"msg": err.Translate(trans)})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		}
		return
	}
	// 2.业务逻辑处理
	if err := logic.UserLogin(u); err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
}
