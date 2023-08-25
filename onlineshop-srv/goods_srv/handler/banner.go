package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"onlineshop-srv/goods_srv/global"
	"onlineshop-srv/goods_srv/model"
	"onlineshop-srv/goods_srv/proto"
)

func (s *GoodsServer) BannerList(_ context.Context, _ *emptypb.Empty) (*proto.BannerListResponse, error) {
	bannerListResponse := proto.BannerListResponse{}

	var banners []model.Banner
	result := global.DB.Find(&banners)

	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerResponses []*proto.BannerResponse
	for _, banner := range banners {
		bannerResponses = append(bannerResponses, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerResponses
	return &bannerListResponse, nil
}
func (s *GoodsServer) CreateBanner(_ context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := model.Banner{}
	banner.Image = req.Image
	banner.Index = req.Index
	banner.Url = req.Url
	global.DB.Save(&banner)

	return &proto.BannerResponse{
		Id: banner.ID,
	}, nil
}
func (s *GoodsServer) DeleteBanner(_ context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "轮播图不存在")
	}

	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateBanner(_ context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	var banner model.Banner
	if result := global.DB.First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "轮播图不存在")
	}
	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}
	global.DB.Save(&banner)
	return &emptypb.Empty{}, nil
}
