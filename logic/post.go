package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/dao/redis"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) error {
	// 生成帖子id
	post.Id = snowflake.GenInt64Id()
	// 保存到redis
	err := redis.CreatePost(fmt.Sprintf("%d", post.Id))
	if err != nil {
		return err
	}
	// 保存到数据库
	return mysql.CreatePost(post)
	// 返回
}

// GetPostDetailById 根据id获取帖子详情
func GetPostDetailById(id int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合数据
	postDetail, err := mysql.GetPostDetailById(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById failed", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}
	// 根据作者id查询作者信息
	username, err := mysql.GetUserById(postDetail.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("user_id", postDetail.AuthorId), zap.Error(err))
		return nil, err
	}
	// 查询社区信息
	communityDetail, err := mysql.GetCommunityDetailById(postDetail.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("comm_id", postDetail.CommunityId), zap.Error(err))
		return nil, err
	}
	// 组合数据
	data = &models.ApiPostDetail{
		AuthorName:      username,
		Post:            postDetail,
		CommunityDetail: communityDetail,
	}
	// 查询数据库中id所对应的数据
	return
}

// GetPostList 获取帖子列表
func GetPostList(offset, limit int64) (data []*models.ApiPostDetail, err error) {
	lists, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(lists))
	for _, post := range lists {
		// 根据作者id查询作者信息
		username, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed", zap.Int64("user_id", post.AuthorId), zap.Error(err))
			continue
		}
		// 查询社区信息
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("comm_id", post.CommunityId), zap.Error(err))
			continue
		}
		// 组合数据
		postDetail := &models.ApiPostDetail{
			AuthorName:      username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postDetail)
	}
	return
}
