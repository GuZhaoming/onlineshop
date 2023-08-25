package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"onlineshop-srv/inventory_srv/proto"
	"sync"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//实例化一个client
	invClient = proto.NewInventoryClient(conn)
}

func main() {
	Init()
	//插入数据
	//var i int32
	//for i = 25; i <= 33; i++ {
	//	TestSetInv(i, 100)
	//}

	//TestInvDetail()

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go TestSell(&wg)
	}
	wg.Wait()

	//TestSell()
	conn.Close()
}

func TestSell(wg *sync.WaitGroup) {
	//并发模拟
	defer wg.Done()
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 30, Num: 1},
		},
	})
	//普通模拟
	//_, err := invClient.Sell(context.Background(), &proto.SellInfo{
	//	GoodsInfo: []*proto.GoodsInvInfo{
	//		//1.扣减成功
	//		//{GoodsId: 26, Num: 10},
	//		//{GoodsId: 27, Num: 10},
	//		//2.库存不足
	//		//{GoodsId: 26, Num: 10},
	//		//{GoodsId: 27, Num: 100},
	//		//3.库存不存在
	//		//{GoodsId: 28, Num: 10},
	//	},
	//})
	if err != nil {
		panic(err)
	}
	fmt.Println("库存扣减成功")
}

func TestSetInv(goodId, num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodId,
		Num:     num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}

func TestInvDetail() {
	rsp, err := invClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: 25,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Num)
}
