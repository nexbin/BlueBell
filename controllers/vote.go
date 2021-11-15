package controllers

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 投票功能
func PostVoteHandler(c *gin.Context) {
	// 1。 获取请求参数（userId.postId）
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果是参数校验错误
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		}
		zap.L().Error("c.ShouldBindJSON failed", zap.Error(err))
		return
	}
	// 2. 记录投票数据
	// 获取用户id
	currentUserId, err := GetCurrentUserId(c)
	if err != nil {
		zap.L().Error("GetCurrentUserId failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	p.UserId = currentUserId
	// 投票
	if err := logic.PostVote(p); err != nil {
		zap.L().Error("logic.PostVote failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
