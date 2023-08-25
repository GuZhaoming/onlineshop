package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	utils "onlineshop-api/goods-web/utils"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"onlineshop-api/goods-web/global"
	"onlineshop-api/goods-web/initialize"
	"onlineshop-api/goods-web/utils/register/consul"
)

func main() {

	//1、初始化logger
	initialize.InitLogger()

	//2.初始化配置信息
	initialize.InitConfig()

	//3.初始化routers
	Router := initialize.Routers()

	//4.初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}

	//5.初始化srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//本地开发端口号固定，线上环境启动获取端口号
	debug := viper.GetBool("onlineshop-debug")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//服务注册
	register_client := consul.NewRegistry(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = register_client.Register("192.168.31.172", global.ServerConfig.ConsulInfo.Port,
		global.ServerConfig.Name, global.ServerConfig.ConsulInfo.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}

	zap.S().Debugf("启动服务器，端口:%d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(":8022"); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
