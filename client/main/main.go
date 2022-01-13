package main

import (
	"chatroom/client/process"
	"fmt"
)

func RoomMenu() ([]byte, error) {
	fmt.Println("---------欢迎登陆多人聊天系统---------")
	fmt.Println("---------1 登陆聊天室---------")
	fmt.Println("---------2 注册用户---------")
	fmt.Println("---------3 退出系统---------")
	fmt.Print("请选择(1-3)：")
	var key int
	var inputData []byte
	// var msg messageMessage
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("---------1 登陆聊天室---------")
		fmt.Println("============================")
		var userId int
		var userPwd string
		fmt.Print("请输入用户的id：")
		fmt.Scanf("%d\n", &userId)
		fmt.Print("请输入用户的密码：")
		fmt.Scanf("%s\n", &userPwd)

		up := &process.UserProcess{}
		up.Login(userId, userPwd)

	case 2:
		fmt.Println("---------2 注册用户---------")
		fmt.Println("============================")

		var userId int
		var userPwd string
		var userName string
		fmt.Print("请输入用户的id：")
		fmt.Scanf("%d\n", &userId)
		fmt.Print("请输入用户的密码：")
		fmt.Scanf("%s\n", &userPwd)
		fmt.Print("请输入用户名：")
		fmt.Scanf("%s\n", &userName)

		up := &process.UserProcess{}
		up.Register(userId, userPwd, userName)

		// err := RoomRegister()

		// if err == nil {
		// 	fmt.Println("注册成功")
		// } else {
		// 	fmt.Println("注册失败，显示错误信息", err)
		// }
	case 3:
		fmt.Println("退出系统")
	default:
		fmt.Println("输入有误，请重新输入")
		return []byte{}, nil
	}

	// inputData, _ := json.Marshal(msg)
	return inputData, nil
}

func main() {

	for {
		// 获取输入的数据
		_, err := RoomMenu()
		if err != nil {
			fmt.Println("输入错误")
		}

	}
}
