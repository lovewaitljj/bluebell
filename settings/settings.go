package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()                  // 读取配置信息
	if err != nil {                             // 读取配置信息失败
		fmt.Printf("viper.readInConfig failed, err:%v\n", err)
		return
	}
	//监控配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了~")
	})
	return

}
