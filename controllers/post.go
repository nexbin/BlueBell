package controllers

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	post := new(models.Post)
	// 1.获取参数和校验
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Error("crete post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从上下文中取出当前发帖的userid
	userId, err := GetCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorId = userId
	// 2.创建帖子
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 查看帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 获取帖子的id（参数）
	idParam := c.Param("id")
	intId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt(post_id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 取出帖子数据
	postDetail, err := logic.GetPostDetailById(intId)
	if err != nil {
		zap.L().Error("logic.GetPostDetailById(intId) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, postDetail)
}

// GetPostList 获取帖子列表
func GetPostList(c *gin.Context) {
	// 从数据库立取出帖子（获取数据）
	// 实现分页
	offset, limit := GetPageSegInfo(c)
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}
