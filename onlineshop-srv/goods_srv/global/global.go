package global

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"onlineshop-srv/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

func init() {
	dsn := "root:root@tcp(192.168.31.172:3306)/olshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //如果一个数据库查询花费的时间超过这个阈值，它将被认为是慢查询
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用结构体名称的单数形式作为表名
		},
	})
	if err != nil {
		panic(err)
	}
}
