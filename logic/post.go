package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
)

func CreatePost(post *models.Post) error {
	// 生成帖子id
	post.Id = snowflake.GenInt64Id()
	// 保存到数据库
	return mysql.CreatePost(post)
	// 返回
}
