package controllers

import (
	"BlueBell/dao/mysql"
	"BlueBell/logic"
	"BlueBell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		} else {
			// 否则为validator.ValidationErrors，返回json格式的新error
			ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		}
		return
	}

	// 2.业务处理
	if err := logic.UserSignUp(p); err != nil {
		zap.L().Error("User sign up failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseErrorWithMsg(c, CodeUserExist, err.Error())
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}

	// 3.返回响应
	ResponseSuccess(c, CodeSuccess.Msg())
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
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Translate(trans))
		} else {
			// 参数不符合要求
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		}
		return
	}
	// 2.业务逻辑处理
	user, err := logic.UserLogin(u)
	if err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		} else if errors.Is(err, mysql.ErrorPwd) {
			ResponseError(c, CodeInvalidPwd)
		} else {
			ResponseErrorWithMsg(c, UnExceptionCode, err.Error())
		}
		return
	}
	// 3.返回响应，返回token
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserId), // 前端能识别的 < (1 << 53) - 1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
