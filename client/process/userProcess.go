package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"fmt"
	"net"
)

type UserProcess struct {
	// Conn net.Conn
}

func (this *UserProcess) Login(userId int, pwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()

	userInfo := message.LoginMes{
		Id:       userId,
		Password: pwd,
	}

	inputData, _ := utils.MessageEncode(message.LoginMesType, userInfo)

	trans := utils.Transfer{
		Conn: conn,
	}
	// fmt.Println("inputData:", string(inputData))

	err = trans.WritePkg(inputData)
	if err != nil {
		return
	}

	var logRes message.LoginResMes
	var resData message.Message
	resData, err = trans.ReadPkg()
	if err != nil {
		return
	}
	err = utils.MessageDecode(&resData, &logRes)
	if logRes.Code == 200 {
		// 初始化curUser
		curUser.Conn = conn
		curUser.Id = userInfo.Id
		curUser.Status = message.UserOffline
		curUser.Name = logRes.UserName

		// 显示在线用户列表
		fmt.Println("当前在线用户列表：")
		for _, id := range logRes.Data {
			if id == userId {
				continue
			}
			user := &message.User{
				Id:     id,
				Status: message.UserOnline,
			}
			onlineUsers[id] = user
			fmt.Println("用户ID：", id)
		}
		fmt.Print("\n\n")

		go serverProcessMes(conn)

		for {
			ShowMenu(&userInfo)
		}
	} else {
		fmt.Println("错误信息：", logRes.Error)
	}
	return
}

func (this *UserProcess) Register(id int, pwd string, Name string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()

	userInfo := &message.RegisterMes{
		User: message.User{
			Id:   id,
			Pwd:  pwd,
			Name: Name,
		},
	}
	inputData, _ := utils.MessageEncode(message.RegisterMesType, userInfo)

	trans := utils.Transfer{
		Conn: conn,
	}

	var resData message.Message
	var resMessage message.ResMessage

	resData, err = trans.WriteReply(inputData)
	if err != nil {
		return
	}
	err = utils.MessageDecode(&resData, &resMessage)
	if resMessage.Code == 200 {
		fmt.Println("注册成功:", userInfo)
		// go serverProcessMes(conn)

		// // for {
		// // 	ShowMenu(&userInfo)
		// // }
		return
	} else {
		fmt.Println("错误信息：", resMessage.Error)
	}
	return
}
