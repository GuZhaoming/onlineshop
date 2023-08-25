package handler

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"onlineshop-srv/inventory_srv/global"
	"onlineshop-srv/inventory_srv/model"
	"onlineshop-srv/inventory_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (i *InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存
	var inv model.Inventory
	//global.DB.First(&inv,req.GoodsId)
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	inv.Goods = req.GoodsId
	inv.Stocks = req.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (i *InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	//库存信息
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}

// 方式一：全局锁
var m sync.Mutex

//
//func (i *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
//	// 扣减库存,目前本地事务
//	//1.假如卖三种，中间一件商品不足，但是前面的已经改到数据库了，(事务)解决
//	//2.假设两个人同时购买，同时提交，数据库同时修改，(加锁)解决
//	tx := global.DB.Begin() //开启事务
//	//方式一、解决普通并发m.Lock()
//	m.Lock() //获取锁，但是，又要考虑性能问题，假如10w人访问同一把锁，GG
//	for _, goodInfo := range req.GoodsInfo {
//		var inv model.Inventory
//		//查询是否存在该商品库存
//		//方式二、使用mysql的悲观锁
//		if result := tx.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
//			tx.Rollback() //回滚事务
//			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
//		}
//		//取出到的数量小于库存
//		if inv.Stocks < goodInfo.Num {
//			tx.Rollback() //回滚事务
//			return nil, status.Errorf(codes.InvalidArgument, "库存不足")
//		}
//		//扣减
//		inv.Stocks -= goodInfo.Num
//		tx.Save(&inv)
//	}
//	tx.Commit() //手动提交事务
//	m.Unlock()  //释放锁:数据库执行完才释放锁
//	return &emptypb.Empty{}, nil
//}

//方式二：悲观锁
//func (i *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
//	tx := global.DB.Begin() //开启事务
//	for _, goodInfo := range req.GoodsInfo {
//		var inv model.Inventory
//		//查询是否存在该商品库存
//		//方式二、使用mysql的悲观锁
//		if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
//			tx.Rollback() //回滚事务
//			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
//		}
//		//取出到的数量小于库存
//		if inv.Stocks < goodInfo.Num {
//			tx.Rollback() //回滚事务
//			return nil, status.Errorf(codes.InvalidArgument, "库存不足")
//		}
//		//扣减
//		inv.Stocks -= goodInfo.Num
//		tx.Save(&inv)
//	}
//	tx.Commit() //手动提交事务
//	return &emptypb.Empty{}, nil
//}

// 方式三、乐观锁
func (i *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := global.DB.Begin() //开启事务
	for _, goodInfo := range req.GoodsInfo {
		fmt.Println(goodInfo)
		var inv model.Inventory
		for {
			//查询是否存在该商品库存
			if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
				tx.Rollback() //回滚事务
				return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
			}
			//判断库存是否充足
			if inv.Stocks < goodInfo.Num {
				tx.Rollback() //回滚事务
				return nil, status.Errorf(codes.InvalidArgument, "库存不足")
			}
			//扣减
			inv.Stocks -= goodInfo.Num
			//update inventory set stocks = stocks-1, version=version+1 where goods=goods and version=version
			if result := tx.Model(model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version = ?", goodInfo.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version + 1}); result.RowsAffected == 0 {
				zap.S().Info("库存扣减失败")
			} else {
				break
			}
		}
	}
	tx.Commit() //手动提交事务
	return &emptypb.Empty{}, nil
}

func (i *InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//库存归还，
	//1.例如加入购物车不买 2.库存扣减成功、但是订单创建失败 3.手动归还
	tx := global.DB.Begin() //开启事务
	m.Lock()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		//查询是否存在该商品库存
		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() //回滚事务
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}
		//增加
		inv.Stocks += goodInfo.Num
		tx.Save(&inv)
	}

	tx.Commit() //手动提交事务
	m.Unlock()  //释放锁:数据库执行完才释放锁
	return &emptypb.Empty{}, nil
}
