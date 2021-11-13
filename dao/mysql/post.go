package mysql

import "BlueBell/models"

func CreatePost(post *models.Post) error {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id
	) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.Id, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}
