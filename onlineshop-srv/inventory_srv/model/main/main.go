package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"onlineshop-srv/inventory_srv/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMd5(code string) string { //md5算法进行加密存储
	Md5 := md5.New()
	io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	dsn := "root:root@tcp(192.168.31.172:3306)/olshop_inventory_srv?charset=utf8mb4&parseTime=True&loc=Local"
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

	_ = db.AutoMigrate(&model.Inventory{}) //将表结构直接生成表

}
