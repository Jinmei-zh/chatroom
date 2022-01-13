package process

import (
	"chatroom/common/message"
	"fmt"
)

// 接收到消息
func recvGroupMes(sms message.SmsMes) {
	fmt.Printf("\n[群消息]用户%d(%s)说：%s\n", sms.Id, sms.Name, sms.Content)
}

// 接收到消息
func recvPrivateMes(sms message.SmsPrivateMes) {
	fmt.Printf("\n[私信]用户%d(%s)说：%s\n", sms.FromUser.Id, sms.FromUser.Name, sms.Content)
}
