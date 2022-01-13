package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
	"os"
)

func ShowMenu(userInfo *message.LoginMes) {
	fmt.Printf("---------欢迎登陆:%d---------\n", userInfo.Id)
	fmt.Println("---------1. 显示用户在线列表---------")
	fmt.Println("---------2. 发送消息---------")
	fmt.Println("---------3. 信息列表---------")
	fmt.Println("---------4. 退出系统---------")
	fmt.Println("---------5. 发送私信---------")

	fmt.Print("请选择(1-5):")
	var key int
	var userId int
	var content string
	smsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1: //显示在线列表
		// fmt.Println("显示用户在线列表")
		outputOnlineUser()
	case 2: //发送消息
		fmt.Println("请输入你想对大家说的话：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroup(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		signOut()
		os.Exit(0)
	case 5:
		fmt.Println("发送私信")
		fmt.Println("请输入发送的用户ID：")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入你想对他说的话：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendUser(content, userId)
	default:
		fmt.Println("输入有误，请重新输入")
	}
}

func serverProcessMes(conn net.Conn) {
	trans := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := trans.ReadPkg()
		if err != nil {
			if err == io.EOF {
				continue
			}
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线/下线了
			var msg message.NotifyUserStatusMes
			err = utils.MessageDecode(&mes, &msg)
			if err != nil {
				return
			}
			updateUserStatus(msg)
		case message.SmsResMesType: //收到群消息
			var sms message.SmsMes
			err = utils.MessageDecode(&mes, &sms)
			if err != nil {
				return
			}
			recvGroupMes(sms)
		case message.SmsPrivateMesType: //收到私聊
			var sms message.SmsPrivateMes
			err = utils.MessageDecode(&mes, &sms)
			if err != nil {
				return
			}
			recvPrivateMes(sms)
		default:
			fmt.Println("服务端返回了未知类型信息：", mes)
		}
	}
}
