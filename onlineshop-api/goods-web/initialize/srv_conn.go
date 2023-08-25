package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"onlineshop-api/goods-web/global"
	"onlineshop-api/goods-web/proto"
)

// 负载均衡
func InitSrvConn2() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】")
	}
	userSrvClient := proto.NewGoodsClient(userConn)
	global.GoodsSrvClient = userSrvClient
}

func InitSrvConn() {
	//cfg := api.DefaultConfig()
	//consulInfo := global.ServerConfig.ConsulInfo
	//cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	//userSrvHost := ""
	//userSrvPort := 0
	//client, err := api.NewClient(cfg)
	//if err != nil {
	//	panic(err)
	//}
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	//if err != nil {
	//	panic(err)
	//}
	//for _, value := range data {
	//	userSrvHost = value.Address
	//	userSrvPort = value.Port
	//	break
	//}
	//if userSrvHost == "" {
	//	zap.S().Fatal("[GetUserList] 连接 用户服务失败")
	//}

	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", "127.0.0.1", 50052), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】")
	}
	//1.后续用户下线。2、改端口。3、改ip，负载均衡来做
	//2、事先建立好连接，不用多次进行tcp的三次握手
	//3、一个goroutine共用，连接池
	userSrvClient := proto.NewGoodsClient(userConn)
	global.GoodsSrvClient = userSrvClient

}
