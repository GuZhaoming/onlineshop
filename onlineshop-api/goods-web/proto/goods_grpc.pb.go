// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: goods.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	// 商品接口
	GoodsList(ctx context.Context, in *GoodsFilterRequest, opts ...grpc.CallOption) (*GoodsListResponse, error)
	BatchGetGoods(ctx context.Context, in *BatchGoodsIdInfo, opts ...grpc.CallOption) (*GoodsListResponse, error)
	CreateGoods(ctx context.Context, in *CreateGoodsInfo, opts ...grpc.CallOption) (*GoodsInfoResponse, error)
	DeleteGoods(ctx context.Context, in *DeleteGoodsInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateGoods(ctx context.Context, in *CreateGoodsInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetGoodsDetail(ctx context.Context, in *GoodInfoRequest, opts ...grpc.CallOption) (*GoodsInfoResponse, error)
	// 商品分类
	GetAllCategorysList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error)
	// 获取子分类
	GetSubCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error)
	CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 品牌和轮播图
	BrandList(ctx context.Context, in *BrandFilterRequest, opts ...grpc.CallOption) (*BrandListResponse, error)
	CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error)
	DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 轮播图
	BannerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BannerListResponse, error)
	CreateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*BannerResponse, error)
	DeleteBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 品牌分类
	CategoryBrandList(ctx context.Context, in *CategoryBrandFilterRequest, opts ...grpc.CallOption) (*CategoryBrandListResponse, error)
	// 通过category获取brands
	GetCategoryBrandList(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*BrandListResponse, error)
	CreateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*CategoryBrandResponse, error)
	DeleteCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GoodsList(ctx context.Context, in *GoodsFilterRequest, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	out := new(GoodsListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) BatchGetGoods(ctx context.Context, in *BatchGoodsIdInfo, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	out := new(GoodsListResponse)
	err := c.cc.Invoke(ctx, "/Goods/BatchGetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateGoods(ctx context.Context, in *CreateGoodsInfo, opts ...grpc.CallOption) (*GoodsInfoResponse, error) {
	out := new(GoodsInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteGoods(ctx context.Context, in *DeleteGoodsInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateGoods(ctx context.Context, in *CreateGoodsInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsDetail(ctx context.Context, in *GoodInfoRequest, opts ...grpc.CallOption) (*GoodsInfoResponse, error) {
	out := new(GoodsInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetAllCategorysList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetAllCategorysList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetSubCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error) {
	out := new(SubCategoryListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error) {
	out := new(CategoryInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) BrandList(ctx context.Context, in *BrandFilterRequest, opts ...grpc.CallOption) (*BrandListResponse, error) {
	out := new(BrandListResponse)
	err := c.cc.Invoke(ctx, "/Goods/BrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error) {
	out := new(BrandInfoResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) BannerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BannerListResponse, error) {
	out := new(BannerListResponse)
	err := c.cc.Invoke(ctx, "/Goods/BannerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*BannerResponse, error) {
	out := new(BannerResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateBanner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CategoryBrandList(ctx context.Context, in *CategoryBrandFilterRequest, opts ...grpc.CallOption) (*CategoryBrandListResponse, error) {
	out := new(CategoryBrandListResponse)
	err := c.cc.Invoke(ctx, "/Goods/CategoryBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetCategoryBrandList(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*BrandListResponse, error) {
	out := new(BrandListResponse)
	err := c.cc.Invoke(ctx, "/Goods/GetCategoryBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*CategoryBrandResponse, error) {
	out := new(CategoryBrandResponse)
	err := c.cc.Invoke(ctx, "/Goods/CreateCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/DeleteCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateCategoryBrand(ctx context.Context, in *CategoryBrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Goods/UpdateCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	// 商品接口
	GoodsList(context.Context, *GoodsFilterRequest) (*GoodsListResponse, error)
	BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
	CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
	DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
	UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
	GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
	// 商品分类
	GetAllCategorysList(context.Context, *emptypb.Empty) (*CategoryListResponse, error)
	// 获取子分类
	GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error)
	CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
	UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
	// 品牌和轮播图
	BrandList(context.Context, *BrandFilterRequest) (*BrandListResponse, error)
	CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error)
	DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
	UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
	// 轮播图
	BannerList(context.Context, *emptypb.Empty) (*BannerListResponse, error)
	CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error)
	DeleteBanner(context.Context, *BannerRequest) (*emptypb.Empty, error)
	UpdateBanner(context.Context, *BannerRequest) (*emptypb.Empty, error)
	// 品牌分类
	CategoryBrandList(context.Context, *CategoryBrandFilterRequest) (*CategoryBrandListResponse, error)
	// 通过category获取brands
	GetCategoryBrandList(context.Context, *CategoryInfoRequest) (*BrandListResponse, error)
	CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error)
	DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
	UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GoodsList(context.Context, *GoodsFilterRequest) (*GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsList not implemented")
}
func (UnimplementedGoodsServer) BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetGoods not implemented")
}
func (UnimplementedGoodsServer) CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedGoodsServer) DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedGoodsServer) UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
}
func (UnimplementedGoodsServer) GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (UnimplementedGoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategorysList not implemented")
}
func (UnimplementedGoodsServer) GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubCategory not implemented")
}
func (UnimplementedGoodsServer) CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedGoodsServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedGoodsServer) UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedGoodsServer) BrandList(context.Context, *BrandFilterRequest) (*BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrandList not implemented")
}
func (UnimplementedGoodsServer) CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedGoodsServer) DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedGoodsServer) UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedGoodsServer) BannerList(context.Context, *emptypb.Empty) (*BannerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BannerList not implemented")
}
func (UnimplementedGoodsServer) CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanner not implemented")
}
func (UnimplementedGoodsServer) DeleteBanner(context.Context, *BannerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (UnimplementedGoodsServer) UpdateBanner(context.Context, *BannerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
func (UnimplementedGoodsServer) CategoryBrandList(context.Context, *CategoryBrandFilterRequest) (*CategoryBrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryBrandList not implemented")
}
func (UnimplementedGoodsServer) GetCategoryBrandList(context.Context, *CategoryInfoRequest) (*BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryBrandList not implemented")
}
func (UnimplementedGoodsServer) CreateCategoryBrand(context.Context, *CategoryBrandRequest) (*CategoryBrandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategoryBrand not implemented")
}
func (UnimplementedGoodsServer) DeleteCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryBrand not implemented")
}
func (UnimplementedGoodsServer) UpdateCategoryBrand(context.Context, *CategoryBrandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategoryBrand not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GoodsList(ctx, req.(*GoodsFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_BatchGetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGoodsIdInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).BatchGetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/BatchGetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).BatchGetGoods(ctx, req.(*BatchGoodsIdInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoods(ctx, req.(*CreateGoodsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteGoods(ctx, req.(*DeleteGoodsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateGoods(ctx, req.(*CreateGoodsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsDetail(ctx, req.(*GoodInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetAllCategorysList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetAllCategorysList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetAllCategorysList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetAllCategorysList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetSubCategory(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateCategory(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateCategory(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_BrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).BrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/BrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).BrandList(ctx, req.(*BrandFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateBrand(ctx, req.(*BrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteBrand(ctx, req.(*BrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateBrand(ctx, req.(*BrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_BannerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).BannerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/BannerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).BannerList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateBanner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteBanner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateBanner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CategoryBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CategoryBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CategoryBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CategoryBrandList(ctx, req.(*CategoryBrandFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetCategoryBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetCategoryBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/GetCategoryBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetCategoryBrandList(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/CreateCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateCategoryBrand(ctx, req.(*CategoryBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/DeleteCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteCategoryBrand(ctx, req.(*CategoryBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Goods/UpdateCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateCategoryBrand(ctx, req.(*CategoryBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoodsList",
			Handler:    _Goods_GoodsList_Handler,
		},
		{
			MethodName: "BatchGetGoods",
			Handler:    _Goods_BatchGetGoods_Handler,
		},
		{
			MethodName: "CreateGoods",
			Handler:    _Goods_CreateGoods_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _Goods_DeleteGoods_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _Goods_UpdateGoods_Handler,
		},
		{
			MethodName: "GetGoodsDetail",
			Handler:    _Goods_GetGoodsDetail_Handler,
		},
		{
			MethodName: "GetAllCategorysList",
			Handler:    _Goods_GetAllCategorysList_Handler,
		},
		{
			MethodName: "GetSubCategory",
			Handler:    _Goods_GetSubCategory_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Goods_CreateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _Goods_DeleteCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _Goods_UpdateCategory_Handler,
		},
		{
			MethodName: "BrandList",
			Handler:    _Goods_BrandList_Handler,
		},
		{
			MethodName: "CreateBrand",
			Handler:    _Goods_CreateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _Goods_DeleteBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _Goods_UpdateBrand_Handler,
		},
		{
			MethodName: "BannerList",
			Handler:    _Goods_BannerList_Handler,
		},
		{
			MethodName: "CreateBanner",
			Handler:    _Goods_CreateBanner_Handler,
		},
		{
			MethodName: "DeleteBanner",
			Handler:    _Goods_DeleteBanner_Handler,
		},
		{
			MethodName: "UpdateBanner",
			Handler:    _Goods_UpdateBanner_Handler,
		},
		{
			MethodName: "CategoryBrandList",
			Handler:    _Goods_CategoryBrandList_Handler,
		},
		{
			MethodName: "GetCategoryBrandList",
			Handler:    _Goods_GetCategoryBrandList_Handler,
		},
		{
			MethodName: "CreateCategoryBrand",
			Handler:    _Goods_CreateCategoryBrand_Handler,
		},
		{
			MethodName: "DeleteCategoryBrand",
			Handler:    _Goods_DeleteCategoryBrand_Handler,
		},
		{
			MethodName: "UpdateCategoryBrand",
			Handler:    _Goods_UpdateCategoryBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}
