package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	router2 "onlineshop-api/goods-web/router"

	"onlineshop-api/goods-web/middlewares"
)

func Routers() *gin.Engine {

	Router := gin.Default()

	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})

	})

	//配置跨域
	Router.Use(middlewares.Cors())
	//配置路由
	ApiGroup := Router.Group("/g/v1")
	router2.InitGoodsRouter(ApiGroup)
	router2.InitCategoryRouter(ApiGroup)
	router2.InitBannerRouter(ApiGroup)
	router2.InitBrandRouter(ApiGroup)
	return Router
}
