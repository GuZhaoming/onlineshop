package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"onlineshop-api/goods-web/api/goods"
	"onlineshop-api/goods-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	zap.S().Info("配置商品相关的URL")
	{
		//商品列表
		GoodsRouter.GET("", goods.List)
		//添加商品
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
		//商品详情
		GoodsRouter.GET("/:id", goods.Detail)
		//删除商品
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)
		//获取库存
		GoodsRouter.GET("/:id/stocks", goods.Stocks)

		//更新商品状态
		GoodsRouter.PATCH("/:id", goods.UpdateStatus)

		//更新商品
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
	}
}
