package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"onlineshop-srv/user_srv/proto"
	"time"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//实例化一个client
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)

		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			PassWord:          "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("qingyu%d", i),
			Mobile:   fmt.Sprintf("1763444905%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserByMobile() {
	res, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "17634449053",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("mobile:%s,nickName:%s,id:%d\n", res.Mobile, res.NickName, res.Id)
}

func TestGetUserById() {
	res, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("mobile:%s,nickName:%s,id:%d\n", res.Mobile, res.NickName, res.Id)
}

func TestUpdateUser() {
	birthday, err := time.Parse("2006-01-02", "2023-08-07")
	if err != nil {
		panic(err)
	}

	res, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       3,
		NickName: "new-qingyu1",
		Birthday: uint64(birthday.Unix()),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}

func main() {
	Init()

	//TestGetUserList()
	//TestCreateUser()
	//TestGetUserByMobile()
	//TestGetUserById()
	TestUpdateUser()
	conn.Close()
}
