package initialize

import (
	"fmt"

	"github.com/spf13/viper"

	"onlineshop-srv/goods_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//读取配置文件
	debug := GetEnvInfo("onlineshop-debug")
	configFileprefix := "config"
	configFileName := fmt.Sprintf("goods_srv/%s-pro.yaml", configFileprefix)
	if debug {
		configFileName = fmt.Sprintf("goods_srv/%s-debug.yaml", configFileprefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
