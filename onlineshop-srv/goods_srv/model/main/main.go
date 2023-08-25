package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"onlineshop-srv/goods_srv/model"
)

func main() {
	dsn := "root:root@tcp(192.168.31.172:3306)/olshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //如果一个数据库查询花费的时间超过这个阈值，它将被认为是慢查询
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用结构体名称的单数形式作为表名
		},
	})
	if err != nil {
		panic(err)
	}

	//_ = db.AutoMigrate(&model.Category{},
	//	&model.Brands{},
	//	&model.GoodsCategoryBrand{},
	//	&model.Banner{},
	//	&model.Goods{}) //将表结构直接生成表

	insertCategoryData(db)
	//insertGoodsCategoryBrandData(db)
	//insertGoodsData(db)
}

func insertCategoryData(db *gorm.DB) {
	categories := []model.Category{
		{Name: "虾", ParentCategoryID: 5, Level: 3, IsTab: false},
		{Name: "鱼", ParentCategoryID: 5, Level: 3, IsTab: false},
		{Name: "螃蟹", ParentCategoryID: 5, Level: 3, IsTab: false},
		{Name: "扇贝", ParentCategoryID: 5, Level: 3, IsTab: false},
		{Name: "鲍鱼", ParentCategoryID: 5, Level: 3, IsTab: false},
		{Name: "鳗鱼", ParentCategoryID: 5, Level: 3, IsTab: false},
		{Name: "章鱼", ParentCategoryID: 5, Level: 3, IsTab: false},
	}

	for _, category := range categories {
		db.Create(&category)
	}
}

func insertGoodsCategoryBrandData(db *gorm.DB) {
	categoryBrandData := []struct {
		CategoryID int
		BrandID    int
	}{
		{CategoryID: 1, BrandID: 6},
		{CategoryID: 1, BrandID: 5},
		{CategoryID: 1, BrandID: 7},
		{CategoryID: 1, BrandID: 8},
		{CategoryID: 2, BrandID: 1},
		{CategoryID: 2, BrandID: 2},
		{CategoryID: 2, BrandID: 3},
		{CategoryID: 2, BrandID: 4},
	}

	for _, data := range categoryBrandData {
		goodsCategoryBrand := model.GoodsCategoryBrand{
			CategoryID: int32(data.CategoryID),
			BrandsID:   int32(data.BrandID),
		}
		if err := db.Create(&goodsCategoryBrand).Error; err != nil {
			panic(err)
		}
	}
}

func insertGoodsData(db *gorm.DB) {
	goods := []model.Goods{
		{
			CategoryID:      1,
			BrandsID:        6,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           true,
			IsHot:           false,
			Name:            "新鲜苹果",
			GoodsSn:         "APP001",
			ClickNum:        100,
			SoldNum:         50,
			FavNum:          20,
			MarketPrice:     2.99,
			ShopPrice:       1.99,
			GoodsBrief:      "新鲜多汁的苹果",
			Images:          model.GormList{"apple1.jpg", "apple2.jpg"},
			DescImages:      model.GormList{"apple_desc1.jpg", "apple_desc2.jpg"},
			GoodsFrontImage: "apple_front.jpg",
		},
		{
			CategoryID:      2,
			BrandsID:        4,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           true,
			IsHot:           false,
			Name:            "新鲜青菜",
			GoodsSn:         "VEG001",
			ClickNum:        80,
			SoldNum:         30,
			FavNum:          15,
			MarketPrice:     1.49,
			ShopPrice:       0.99,
			GoodsBrief:      "绿叶嫩香的青菜",
			Images:          model.GormList{"vegetable1.jpg", "vegetable2.jpg"},
			DescImages:      model.GormList{"vegetable_desc1.jpg", "vegetable_desc2.jpg"},
			GoodsFrontImage: "vegetable_front.jpg",
		}, {
			CategoryID:      1,
			BrandsID:        6,
			OnSale:          true,
			ShipFree:        false,
			IsNew:           false,
			IsHot:           true,
			Name:            "新鲜橙子",
			GoodsSn:         "OR001",
			ClickNum:        120,
			SoldNum:         70,
			FavNum:          25,
			MarketPrice:     2.49,
			ShopPrice:       1.79,
			GoodsBrief:      "酸甜可口的橙子",
			Images:          model.GormList{"orange1.jpg", "orange2.jpg"},
			DescImages:      model.GormList{"orange_desc1.jpg", "orange_desc2.jpg"},
			GoodsFrontImage: "orange_front.jpg",
		},
		{
			CategoryID:      2,
			BrandsID:        1,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           false,
			IsHot:           false,
			Name:            "新鲜鸡胸肉",
			GoodsSn:         "CHX001",
			ClickNum:        90,
			SoldNum:         40,
			FavNum:          18,
			MarketPrice:     8.99,
			ShopPrice:       6.99,
			GoodsBrief:      "嫩滑美味的鸡胸肉",
			Images:          model.GormList{"chicken_breast1.jpg", "chicken_breast2.jpg"},
			DescImages:      model.GormList{"chicken_breast_desc1.jpg", "chicken_breast_desc2.jpg"},
			GoodsFrontImage: "chicken_breast_front.jpg",
		}, {
			CategoryID:      1,
			BrandsID:        6,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           true,
			IsHot:           false,
			Name:            "新鲜草莓",
			GoodsSn:         "STR001",
			ClickNum:        110,
			SoldNum:         60,
			FavNum:          22,
			MarketPrice:     4.29,
			ShopPrice:       3.49,
			GoodsBrief:      "鲜红多汁的草莓",
			Images:          model.GormList{"strawberry1.jpg", "strawberry2.jpg"},
			DescImages:      model.GormList{"strawberry_desc1.jpg", "strawberry_desc2.jpg"},
			GoodsFrontImage: "strawberry_front.jpg",
		},
		{
			CategoryID:      2,
			BrandsID:        5,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           true,
			IsHot:           false,
			Name:            "天然矿泉水",
			GoodsSn:         "WTR001",
			ClickNum:        150,
			SoldNum:         80,
			FavNum:          30,
			MarketPrice:     1.99,
			ShopPrice:       1.49,
			GoodsBrief:      "清凉爽口的矿泉水",
			Images:          model.GormList{"water1.jpg", "water2.jpg"},
			DescImages:      model.GormList{"water_desc1.jpg", "water_desc2.jpg"},
			GoodsFrontImage: "water_front.jpg",
		}, {
			CategoryID:      2,
			BrandsID:        1,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           true,
			IsHot:           false,
			Name:            "新鲜牛肉",
			GoodsSn:         "BEEF001",
			ClickNum:        120,
			SoldNum:         70,
			FavNum:          25,
			MarketPrice:     12.99,
			ShopPrice:       9.99,
			GoodsBrief:      "肉质鲜嫩的牛肉",
			Images:          model.GormList{"beef1.jpg", "beef2.jpg"},
			DescImages:      model.GormList{"beef_desc1.jpg", "beef_desc2.jpg"},
			GoodsFrontImage: "beef_front.jpg",
		},
		{
			CategoryID:      1,
			BrandsID:        6,
			OnSale:          true,
			ShipFree:        true,
			IsNew:           true,
			IsHot:           false,
			Name:            "新鲜樱桃",
			GoodsSn:         "CHY001",
			ClickNum:        90,
			SoldNum:         40,
			FavNum:          18,
			MarketPrice:     6.99,
			ShopPrice:       4.99,
			GoodsBrief:      "甜美多汁的樱桃",
			Images:          model.GormList{"cherry1.jpg", "cherry2.jpg"},
			DescImages:      model.GormList{"cherry_desc1.jpg", "cherry_desc2.jpg"},
			GoodsFrontImage: "cherry_front.jpg",
		},
	}

	for _, good := range goods {
		db.Create(&good)
	}
}
