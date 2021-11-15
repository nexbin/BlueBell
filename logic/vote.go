package logic

import (
	"BlueBell/dao/redis"
	"BlueBell/models"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 3600 * 24 * 7
	scorePreVote     = 432
)

var (
	ErrorVoteTimeExpired = errors.New("投票时间已过期")
)

// PostVote 本项目使用简化版的投票算法
// 投一票 +432 86400/200->需要200张赞成票可以给你的帖子续一天
// 为帖子投票的函数
/*
case1 : 之前没投过票，新增投票          d:0       abs +1/-1     + 432 / - 432
case2 : 之前投反对票，改投赞成票/取消    d:-1      abs 2/1       + 432 * 2 / + 432
case3 : 之前投赞成票，改投反对票/取消    d:1       abs -2/-1     - 432 * 2 / - 432

投票限制: 超过指定时间就不允许再投票了
		到期之后将redis中保存的票数存储在mysql中
		并删除KeyPostVotedPrefix
*/
func PostVote(postParam *models.ParamVoteData) (err error) {
	// 1. 判断投票限制(取帖子发布时间)
	postId := fmt.Sprintf("%d", postParam.PostId)
	userId := fmt.Sprintf("%d", postParam.UserId)
	postTime, err := redis.GetPostReleaseTimeById(postId)
	if err != nil {
		zap.L().Error("get post time failed", zap.Error(err))
		return err
	}
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpired
	}
	// 2. 查看该用户在该帖子的投票状态
	voteStatus, err := redis.CheckUserVotingStatus(postId, userId)
	if err != nil {
		zap.L().Error("check user voting status failed", zap.Error(err))
		return err
	}
	// 3. 更新帖子的分数和用户投票数据
	var dir float64
	if float64(postParam.Direction) > voteStatus {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(float64(postParam.Direction) - voteStatus)
	score := dir * diff * scorePreVote
	if err = redis.UpdateVoteInfo(postId, userId, score, float64(postParam.Direction)); err != nil {
		zap.L().Error("update Vote Info failed", zap.Error(err))
	}
	return
}
