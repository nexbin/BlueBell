package controllers

import (
	"BlueBell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 查询社区分类
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区分类, 以列表的形式返回
	dataList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露在外
		return
	}
	//fmt.Println(dataList)
	ResponseSuccess(c, dataList)
}

// CommunityDetailHandler 根据社区id获取详细信息
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	cId := c.Param("id")
	id, err := strconv.ParseInt(cId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id查询社区信息
	details, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, details)
}
