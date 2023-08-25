package handler

import (
	"context"

	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"onlineshop-srv/goods_srv/global"
	"onlineshop-srv/goods_srv/model"
	"onlineshop-srv/goods_srv/proto"
)

type Server struct {
	proto.UnimplementedGoodsServer
}

// GetAllCategorysList 获取全部分类
func (s *GoodsServer) GetAllCategorysList(_ context.Context, _ *emptypb.Empty) (*proto.CategoryListResponse, error) {

	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)

	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

// GetSubCategory 获取子分类
func (s *GoodsServer) GetSubCategory(_ context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	//结果对象
	categoryListResponse := proto.SubCategoryListResponse{}

	//查询分类是否存在
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "商品分类不存在")
	}

	//结果一
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	//结果二
	var subcategorys []model.Category
	var subCategorysResponse []*proto.CategoryInfoResponse
	//preloads := "SubCategory"
	//if category.Level == 1 {
	//	preloads = "SubCategory.SubCategory"
	//}
	global.DB.Where(&model.Category{Level: req.Id}).Find(&subcategorys)

	for _, subCategory := range subcategorys {
		subCategorysResponse = append(subCategorysResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListResponse.SubCategorys = subCategorysResponse
	return &categoryListResponse, nil
}

func (s *GoodsServer) CreateCategory(_ context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}
	category.Name = req.Name
	category.Level = req.Level

	if req.Level != 1 {
		category.ParentCategoryID = req.ParentCategory
	}

	global.DB.Save(&category)

	return &proto.CategoryInfoResponse{Id: category.ID}, nil
}
func (s *GoodsServer) DeleteCategory(_ context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "商品分类不存在")
	}

	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateCategory(_ context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "商品分类不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}

	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}

	if req.Level != 0 {
		category.Level = req.Level
	}

	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
