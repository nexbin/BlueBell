package mysql

import "BlueBell/models"

func CreatePost(post *models.Post) error {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id
	) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.Id, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}

func GetPostDetailById(pid int64) (data *models.Post, err error) {
	postDetail := new(models.Post)
	sqlStr := `select title,content,post_id,author_id,community_id,create_time,update_time,status from post where post_id = ?`
	err = db.Get(postDetail, sqlStr, pid)
	return postDetail, err
}

// GetPostList 获取Post列表
func GetPostList(offset, limit int64) (posts []*models.Post, err error) {
	// 限制只取两条
	sqlStr := `select title,content,post_id,author_id,community_id,create_time,update_time,status from post limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (offset-1)*limit, limit)
	return
}
