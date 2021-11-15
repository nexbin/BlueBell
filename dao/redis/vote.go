package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func GetPostReleaseTimeById(postId string) (releaseTime float64, err error) {
	//根据帖子id获取帖子发布时间
	releaseTime, err = rdb.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postId).Result()
	if err != nil {
		fmt.Println(releaseTime, err)
		return 0, ErrorGetPostReleaseTime
	}
	return
}

// CheckUserVotingStatus 查看用户在该帖子的投票状态
func CheckUserVotingStatus(postId, userId string) (votingStatus float64, err error) {
	votingStatus, err = rdb.ZScore(ctx, getRedisKey(KeyPostVotedPrefix+postId), userId).Result()
	if err == redis.Nil {
		votingStatus = 0
	} else if err != nil {
		return votingStatus, ErrorCheckUserVotingStatus
	}
	return votingStatus, nil
}

// UpdateVoteInfo 更新帖子分数和用户投票数据
func UpdateVoteInfo(postId string, userId string, score float64, curVoteStatus float64) (err error) {
	// 开启事务
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), score, postId)
	if curVoteStatus == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedPrefix+postId), userId)
	}
	pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedPrefix+postId), &redis.Z{
		Score:  curVoteStatus, // 当前用户投的是赞成票还是反对票
		Member: userId,
	})
	_, err = pipeline.Exec(ctx)
	return
}
