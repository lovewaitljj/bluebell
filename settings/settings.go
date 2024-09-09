package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") // 指定配置文件路径
	// 读取配置信息
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.readInConfig failed, err:%v\n", err)
		return
	}
	if err = viper.Unmarshal(Conf); err != nil {
		return
	}
	//监控配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了~")
		if err = viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper.Unmarshal err", zap.String("err", err.Error()))
		}
	})
	return

}
