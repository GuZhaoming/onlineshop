package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"onlineshop-srv/user_srv/model"
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
	dsn := "root:root@tcp(192.168.31.172:3306)/olshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
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
	//
	//options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(newPassword)
	//
	//for i := 0; i < 10; i++ {
	//	user := model.User{
	//		NickName: fmt.Sprintf("qingyu%d", i),
	//		Mobile:   fmt.Sprintf("1763444905%d", i),
	//		PassWord: newPassword,
	//	}
	//	db.Save(&user)
	//}

	_ = db.AutoMigrate(&model.User{}) //将表结构直接生成表
	//
	////实际开发：盐值长度 迭代次数 key长度 可以使用方法替代
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("generic password", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(newPassword)
	//fmt.Println(len(newPassword))
	//
	//passwordInfo := strings.Split(newPassword, "$")
	//fmt.Println(passwordInfo)
	//check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check) // true

}
