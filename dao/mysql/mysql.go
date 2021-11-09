package mysql

import (
	"BlueBell/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(config *settings.MysqlConfig) (err error) {
	// 也可以使用 sqlx.MustConnect()
	db, err = sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DbName,
		),
	)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conn")) // 设置最大连接数
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn")) // 设置最大闲置连接数
	return
}

func Close() {
	_ = db.Close()
}
