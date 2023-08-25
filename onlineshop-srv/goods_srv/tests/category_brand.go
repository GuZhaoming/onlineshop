package main

import (
	"context"
	"fmt"

	"onlineshop-srv/goods_srv/proto"
)

func GetCategoryBrandList() {
	rsp, err := goodsClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Data)
}
