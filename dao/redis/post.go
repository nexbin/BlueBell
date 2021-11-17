package redis

import (
	"BlueBell/models"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func CreatePost(postId string, communityId int64) (err error) {
	// 启动事务
	pipeline := rdb.TxPipeline()
	// 添加帖子时间
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	// 添加帖子分数
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  0,
		Member: postId,
	})
	// 把帖子加入对应的社区
	cKey := getRedisKey(KeyCommunitySetPrefix) + strconv.Itoa(int(communityId))
	pipeline.SAdd(ctx, cKey, postId)
	_, err = pipeline.Exec(ctx)
	return
}

func getIdsFromKey(key string, offset, limit int64) ([]string, error) {
	start := (offset - 1) * limit
	end := start + limit - 1
	return rdb.ZRevRange(ctx, key, start, end).Result()
}

// GetPostIdsByOrder 根据order获取帖子ids
func GetPostIdsByOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIdsFromKey(key, p.Offset, p.Limit)
}

// CheckVoteInfoByIds 批量返回帖子的赞成票投票情况
func CheckVoteInfoByIds(ids []string) (voteInfos []int64, err error) {
	voteInfos = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedPrefix) + id
	//	// 查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
	//	val1 := rdb.ZCount(ctx, key, "1", "1").Val()
	//	voteInfos = append(voteInfos, val1)
	//}

	// 使用pipeline一次发送多条命令，减少RTT
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedPrefix) + id
		pipeline.ZCount(ctx, key, "1", "1")
	}
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		voteInfos = append(voteInfos, v)
	}
	return
}

func GetCommunityPostListById(p *models.ParamPostList) (ids []string, err error) {
	// 使用zinterstore 把分区的帖子set与帖子分数zset生成一个新的zset
	// 针对新的zset按之前的逻辑取数据
	// 社区的key
	ckey := getRedisKey(KeyCommunitySetPrefix) + fmt.Sprintf("%d", p.CommunityId)
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + fmt.Sprintf("%d", p.CommunityId)
	if rdb.Exists(ctx, key).Val() < 1 {
		pipeline := rdb.Pipeline()
		// 如果这个key不存在
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{ckey, orderKey},
			Weights:   nil,
			Aggregate: "MAX",
		}) // zinterstore计算
		pipeline.Expire(ctx, key, 60*time.Second) // 设置超时时间
		_, err = pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	// 存在缓存就直接查询
	return getIdsFromKey(key, p.Offset, p.Limit)
}
