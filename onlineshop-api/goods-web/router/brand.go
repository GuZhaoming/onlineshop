package router

import (
	"github.com/gin-gonic/gin"

	"onlineshop-api/goods-web/api/brands"
)

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brands")
	{
		// 品牌列表页
		BrandRouter.GET("", brands.BrandList)
		// 删除品牌
		BrandRouter.DELETE("/:id", brands.DeleteBrand)
		//新建品牌
		BrandRouter.POST("", brands.NewBrand)
		//修改品牌信息
		BrandRouter.PUT("/:id", brands.UpdateBrand)
	}

	CategoryBrandRouter := Router.Group("categorybrands")
	{
		CategoryBrandRouter.GET("", brands.CategoryBrandList)          // 类别品牌列表页
		CategoryBrandRouter.DELETE("/:id", brands.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("", brands.NewCategoryBrand)          //新建类别品牌
		CategoryBrandRouter.PUT("/:id", brands.UpdateCategoryBrand)    //修改类别品牌
		CategoryBrandRouter.GET("/:id", brands.GetCategoryBrandList)   //获取分类的品牌
	}
}
