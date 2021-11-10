package mysql

import (
	"BlueBell/models"
	"crypto/md5"
	"encoding/hex"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务进行调用

// CheckUserExist 检查指定用户的用户名是否存在
func CheckUserExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(u *models.User) (err error) {
	// 对新用户的密码进行加密
	u.Password = encryptPwd(u.Password)
	// 执行SQL入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, u.UserId, u.Username, u.Password)
	return
}

func encryptPwd(pwd string) string {
	h := md5.New()
	h.Write([]byte("nexbin.com")) // 加盐字符串
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}
