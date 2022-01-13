package processes

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"net"
)

// 消息模块
type SmsProcess struct {
}

func (this *SmsProcess) SendGroup(mes *message.Message) (err error) {
	// 遍历在线用户，排除自己之外的用户发送消息
	var sms message.SmsMes
	err = utils.MessageDecode(mes, &sms)
	if err != nil {
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == sms.Id {
			continue
		}

		// 组装消息
		var data []byte
		data, err = utils.MessageEncode(message.SmsResMesType, sms)
		if err != nil {
			return
		}

		this.SendMesToUser(data, up.Conn)
	}
	return
}

func (this *SmsProcess) SendUserId(mes *message.Message) (err error) {

	var sms message.SmsPrivateMes
	err = utils.MessageDecode(mes, &sms)
	if err != nil {
		return
	}

	// 转发前获取用户昵称
	user, ok := userMgr.onlineUsers[sms.ToUserId]
	if !ok {
		// 没有这个在线用户
		err = model.ERROR_USER_NOTEXISTS
		return
	}

	// 组装消息
	var data []byte
	data, err = utils.MessageEncode(message.SmsPrivateMesType, sms)
	if err != nil {
		return
	}
	return this.SendMesToUser(data, user.Conn)
}

func (this *SmsProcess) SendMesToUser(data []byte, conn net.Conn) (err error) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}
	return
}
