package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
)

type SmsProcess struct {
}

// 群发消息
func (this *SmsProcess) SendGroup(content string) (err error) {
	var mes message.SmsMes
	mes.Content = content
	mes.Id = curUser.Id
	mes.Status = curUser.Status

	data, err := utils.MessageEncode(message.SmsMesType, mes)
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	return
}

//私信
func (this *SmsProcess) SendUser(content string, userId int) (err error) {
	var mes message.SmsPrivateMes
	mes.Content = content
	mes.FromUser.Id = curUser.Id
	mes.FromUser.Status = curUser.Status
	mes.FromUser.Name = curUser.Name
	mes.ToUserId = userId

	data, err := utils.MessageEncode(message.SmsPrivateMesType, mes)
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	return
}
