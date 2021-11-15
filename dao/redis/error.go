package redis

import "errors"

var (
	ErrorGetPostReleaseTime    = errors.New("获取帖子时间失败")
	ErrorCheckUserVotingStatus = errors.New("查看用户投票状态失败")
	ErrorUpdatePostScore       = errors.New("更新帖子分数失败")
	ErrorUpdateUserVoteInfo    = errors.New("更新用户投票信息失败")
)
