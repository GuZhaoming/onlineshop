package main

import (
	"context"
	"fmt"

	"onlineshop-srv/goods_srv/proto"
)

func TestGoodsList() {
	rsp, err := goodsClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 1,
		PriceMin:    4,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func BatchGetGoods() {
	rsp, err := goodsClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{25, 27},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)

	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func TestGetGoodsDetail() {
	rsp, err := goodsClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 25,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name)

}
