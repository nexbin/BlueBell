package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var GlobalConfig *AppConfig

type AppConfig struct {
	AppName      string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineId    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level     string `mapstructure:"level"`
	FileName  string `mapstructure:"filename"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxAge    int    `mapstructure:"max_age"`
	MaxBackup int    `mapstructure:"max_backup"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxOpenConn string `mapstructure:"max_open_conn"`
	MaxIdleConn string `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile(`./config/config.yaml`)
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		// 读取配置失败
		fmt.Printf("viper load config failed: %v\n", err)
		return err
	}

	// 把读取到的配置反序列化
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Printf("viper.Unmarshal() failed: %v\n", err)
	}

	// 支持热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			zap.L().Error("viper.Unmarshal() failed", zap.Error(err))
		}
		fmt.Printf("%s changed..\n", in.Name)
		fmt.Println(GlobalConfig.MysqlConfig.DbName)
	})
	return
}
