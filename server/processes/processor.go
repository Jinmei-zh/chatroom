package processes

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		user := &UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProcessLogin(mes)
	case message.RegisterMesType:
		user := &UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProcessRegster(mes)
	case message.SignOutMesType: //退出登陆
		user := &UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProcessLoginOut(mes)
	case message.SmsMesType: //群发消息
		sms := &SmsProcess{}
		sms.SendGroup(mes)
	case message.SmsPrivateMesType: //私信消息
		sms := &SmsProcess{}
		sms.SendUserId(mes)
	default:
		fmt.Println("没有处理，接收到数据", mes)
		return
	}
	return
}

func (this *Processor) MainProcess() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		var mes message.Message
		mes, err = tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("断开连接")
			} else {
				fmt.Println("读取长度数据错误 err=", err)
			}
			return
		}
		err = this.ServerProcessMes(&mes)

		if err != nil {
			return
		}

	}
}
