package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
)

// GetCommunityList 获取社区分类
func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，找到community并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetail 获取社区具体信息
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailById(id)
}
