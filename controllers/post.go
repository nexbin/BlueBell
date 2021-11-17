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

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
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

// GetPostListHandler2 获取帖子列表升级版
// 按时间或者分数排序
// 根据前端传来的参数（按分数、按创建时间 排序）动态获取帖子列表
// 1. 获取参数
// 2. 去redis查询id列表
// 3. 根据id去数据库查询详细信息
// 4. 返回响应
func GetPostListHandler2(c *gin.Context) {
	// 1. 获取请求参数
	p := &models.ParamPostList{
		Offset: 1,
		Limit:  10,
		Order:  models.OrderTime,
	} // 默认值
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list handler2 with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 数据查询
	list, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, list)
}

// CommunityPostListHandler 获取当前社区对应的帖子信息
//func CommunityPostListHandler(c *gin.Context) {
//	// 1. 获取请求参数
//	p := &models.ParamCommunityPostList{
//		CommunityId: 0,
//		ParamPostList: &models.ParamPostList{
//			Offset: 1,
//			Limit:  10,
//			Order:  models.OrderTime,
//		},
//	} // 默认值
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("get community post list handler with invalid param")
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	// 2. 数据查询
//	list, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetCommunityPostList failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	// 3.返回响应
//	ResponseSuccess(c, list)
//}
