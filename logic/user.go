package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"errors"
	"go.uber.org/zap"
)

func UserSignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	if isExist, err := mysql.CheckUserExist(p.Username); isExist {
		return errors.New("用户已存在")
	} else if err != nil {
		// 数据库查询出错
		zap.L().Error("Check user exist failed,", zap.Error(err))
		return err
	}
	// 2. 生成UID
	userId := snowflake.GenInt64Id()
	// 3.构造user示例
	u := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 4. 数据入库
	if err := mysql.InsertUser(u); err != nil {
		zap.L().Error("Insert userdata into mysql failed", zap.Error(err))
		return err
	}
	return
}