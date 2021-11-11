package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/jwt"
	"BlueBell/pkg/snowflake"
	"go.uber.org/zap"
)

// UserSignUp 用户注册逻辑
func UserSignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
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

// UserLogin 用户登录逻辑
func UserLogin(u *models.ParamLogin) (string, error) {
	user := &models.User{
		Username: u.Username,
		Password: u.Password,
	}
	// 传递的是指针，可以在里面取到userid
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成JWT
	return jwt.GenToken(user.UserId, user.Username)
}
