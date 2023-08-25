package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	utils "onlineshop-api/user-web/utils"
	myvalidator "onlineshop-api/user-web/validator"

	"onlineshop-api/user-web/global"
	"onlineshop-api/user-web/initialize"
	"onlineshop-api/user-web/utils/register/consul"
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

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
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
		if err := Router.Run(":8021"); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
