package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"onlineshop-api/oss-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//1.读取本地配置
	debug := GetEnvInfo("onlineshop-debug")
	configFileprefix := "config"
	configFileName := fmt.Sprintf("goods-web/%s-pro.yaml", configFileprefix)
	if debug {
		configFileName = fmt.Sprintf("goods-web/%s-debug.yaml", configFileprefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息:%v", global.ServerConfig)
	//2.动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件发生变化%v", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息:%v", global.ServerConfig)
	})
}
