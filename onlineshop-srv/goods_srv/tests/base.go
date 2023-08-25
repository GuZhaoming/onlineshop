package main

import (
	"google.golang.org/grpc"

	"onlineshop-srv/goods_srv/proto"
)

var goodsClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//实例化一个client
	goodsClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	//TestGetGoodsList()
	//TestGetAllCategoryList()
	//TestGetSubCategory()
	//GetCategoryBrandList()
	//TestGoodsList()
	//BatchGetGoods()
	TestGetGoodsDetail()
	conn.Close()

}
