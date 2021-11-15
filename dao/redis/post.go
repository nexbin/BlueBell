package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

func CreatePost(postId string) (err error) {
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
	_, err = pipeline.Exec(ctx)
	return
}
