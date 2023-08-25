package main

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"onlineshop-srv/goods_srv/proto"
)

func TestGetAllCategoryList() {
	rsp, err := goodsClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)

}

func TestGetSubCategory() {
	rsp, err := goodsClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 1,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
}
