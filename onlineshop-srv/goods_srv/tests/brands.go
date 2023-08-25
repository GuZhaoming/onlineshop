package main

import (
	"context"
	"fmt"

	"onlineshop-srv/goods_srv/proto"
)

func TestGetGoodsList() {
	rsp, err := goodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}
