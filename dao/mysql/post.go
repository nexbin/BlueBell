package mysql

import (
	"BlueBell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

func CreatePost(post *models.Post) error {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id
	) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.Id, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}

// GetPostDetailById 根据ID查询单个帖子详细信息
func GetPostDetailById(pid int64) (data *models.Post, err error) {
	postDetail := new(models.Post)
	sqlStr := `select title,content,post_id,author_id,community_id,create_time,update_time,status from post where post_id = ?`
	err = db.Get(postDetail, sqlStr, pid)
	return postDetail, err
}

// GetPostList 获取Post列表
func GetPostList(offset, limit int64) (posts []*models.Post, err error) {
	// 限制只取两条
	sqlStr := `select title,content,post_id,author_id,community_id,create_time,update_time,status from post
	order by create_time DESC
	limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (offset-1)*limit, limit)
	return
}

// GetPostListByIds 根据给定的ID列表查询数据
func GetPostListByIds(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	db.Rebind(query)
	postList = make([]*models.Post, 0, len(ids))
	err = db.Select(&postList, query, args...)
	return
}
