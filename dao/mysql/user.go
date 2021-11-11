package mysql

import (
	"BlueBell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
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

// encryptPwd 密码加密
func encryptPwd(pwd string) string {
	h := md5.New()
	h.Write([]byte("nexbin.com")) // 加盐字符串
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

// Login 查找username所对应的password
func Login(u *models.User) error {
	// 将用户名所对应的信息从数据库中取出
	oPwd := u.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err := db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return errors.New(fmt.Sprintf("%s:该用户不存在", u.Username))
	} else if err != nil {
		return err
	}
	// 数据库中的密码和输入的密码进行比对
	if encryptPwd(oPwd) != u.Password {
		return errors.New("登录失败，密码不一致")
	}
	return nil
}
