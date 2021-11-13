package mysql

import (
	"BlueBell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (list []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err = db.Select(&list, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}

	return
}

func GetCommunityDetailById(id int64) (cd *models.CommunityDetail, err error) {
	cd = new(models.CommunityDetail)
	sqlStr := `select community_id, 
				community_name, 
				introduction, 
				create_time, 
				update_time 
				from community 
				where id = ?`
	if err = db.Get(cd, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidId
		}
	}
	return
}
