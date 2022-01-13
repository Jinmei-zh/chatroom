package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"chatroom/server/utils"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUser model.CurUser //登陆成功后初始化

// 退出登陆
func signOut() (err error) {
	signout := &message.SignOutMes{
		UserId: curUser.Id,
	}
	data, err := utils.MessageEncode(message.SignOutMesType, signout)
	if err != nil {
		return
	}
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	return tf.WritePkg(data)
}

func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")

	statusItem := map[int]string{
		message.UserOnline:  "在线",
		message.UserOffline: "离线",
	}

	for id, v := range onlineUsers {
		fmt.Printf("用户id:\t%d [%s]\n", id, statusItem[v.Status])
	}
}

func updateUserStatus(msg message.NotifyUserStatusMes) {
	user, ok := onlineUsers[msg.UserId]
	if !ok {
		user = &message.User{
			Id: msg.UserId,
		}
	}
	user.Status = msg.Status //更改状态
	onlineUsers[msg.UserId] = user

	// 上线提示消息
	fmt.Printf("\n[新消息]用户%d(%s)上线了\n", user.Id, user.Name)

	outputOnlineUser()
}
