package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrorUserNotLogin = errors.New("用户未登录")

	ContextUserIdKey = "userId"
)

// GetCurrentUserId 根据用户id获取用户信息
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

// GetPageSegInfo 获取分页信息
func GetPageSegInfo(c *gin.Context) (offset, limit int64) {
	var err error
	offsetStr := c.Query("offset")
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}
	limitStr := c.Query("limit")
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
