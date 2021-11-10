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
		errs, ok := err.(validator.ValidationErrors)
		// 如果错误类型不为validator.ValidationErrors类型，说明可能是其他错误,比如序列化错误等
		if !ok {
			c.JSON(http.StatusOK, &gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			// 否则为validator.ValidationErrors，返回json格式的新error
			c.JSON(http.StatusOK, &gin.H{
				"msg": errs.Translate(trans),
			})
		}
		zap.L().Error("Sign up with invalid param,", zap.Error(err))
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
