package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotLogin = errors.New("用户未登录")
	ContextUserIdKey  = "userId"
)

func GetCurrentUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get(ContextUserIdKey)
	if !ok {
		return 0, ErrorUserNotLogin
	}
	userId, ok := id.(int64)
	if !ok {
		return 0, ErrorUserNotLogin
	}
	return userId, nil
}
