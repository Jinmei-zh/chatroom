package processes

// 用户登陆模块
import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int, status int) {
	for id, up := range userMgr.onlineUsers {
		fmt.Printf("id=%d,up=%v\n", id, up)
		// 过滤自己
		if id == userId {
			continue
		}
		// 通知其他人我上线了
		up.NotifyToOther(userId, status)
	}
}

func (this *UserProcess) NotifyToOther(userId int, status int) (err error) {
	// 通知其他用户新用户上线
	userStatus := &message.NotifyUserStatusMes{
		UserId: userId,
		Status: status,
	}

	msg, err := utils.MessageEncode(message.NotifyUserStatusMesType, userStatus)
	if err != nil {
		return
	}

	//发送
	trans := &utils.Transfer{
		Conn: this.Conn,
	}
	trans.WritePkg(msg)
	return
}

func (this *UserProcess) ServerProcessRegster(mes *message.Message) (err error) {
	var data message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &data)
	if err != nil {
		fmt.Println(" json.Unmarshal fail err=", err)
		return
	}

	// 返回数据
	resMes := message.LoginResMes{
		Code:  200,
		Error: "",
	}

	err = model.MyUserDao.Register(data.User.Id, data.User.Pwd, data.User.Name)
	if err != nil {
		resMes.Error = err.Error()
		if err == model.ERROR_USER_NOTEXISTS {
			resMes.Code = 402
		} else if err == model.ERROR_USER_PWD {
			resMes.Code = 403
		} else {
			resMes.Code = 500
			resMes.Error = "服务器内部错误"
		}
	} else {
		fmt.Println("注册成功", data.User)
	}
	fmt.Println("返回提示信息：", resMes)

	resData, err := json.Marshal(resMes)

	message := message.Message{
		Type: message.RegisterResMesType,
		Data: string(resData),
	}

	msg, err := json.Marshal(message)

	// 发送
	trans := &utils.Transfer{
		Conn: this.Conn,
	}
	trans.WritePkg(msg)

	return
}

// 退出登陆
func (this *UserProcess) ServerProcessLoginOut(mes *message.Message) (err error) {
	var revData message.SignOutMes
	err = utils.MessageDecode(mes, &revData)
	if err != nil {
		return
	}

	// 全局在线用户去除
	userMgr.DelOnlineUser(revData.UserId)

	//通知其他用户我下线了
	this.NotifyOthersOnlineUser(revData.UserId, message.UserOffline)

	return
}

// 登陆验证
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var logMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &logMes)
	if err != nil {
		fmt.Println(" json.Unmarshal fail err=", err)
		return
	}

	// 返回数据
	resMes := &message.LoginResMes{
		Code:  200,
		Error: "",
	}

	user, err := model.MyUserDao.Login(logMes.Id, logMes.Password)
	if err != nil {
		resMes.Error = err.Error()
		if err == model.ERROR_USER_NOTEXISTS {
			resMes.Code = 402
		} else if err == model.ERROR_USER_PWD {
			resMes.Code = 403
		} else {
			resMes.Code = 500
			resMes.Error = "服务器内部错误"
		}
	} else {
		fmt.Println("用户登陆信息：", user)

		// 登陆成功，记录到在线用户
		this.UserId = logMes.Id //登陆成功后记录用户id信息
		userMgr.AddOnlineUser(this)

		// 获取在线用户列表
		resMes.Data = userMgr.GetAllOnlineUserIds()
		resMes.UserName = user.Name

		//通知其他用户上线状态
		this.NotifyOthersOnlineUser(logMes.Id, message.UserOnline)

	}
	fmt.Println("返回提示信息：", resMes)

	trans := &utils.Transfer{
		Conn: this.Conn,
	}
	resData, err := json.Marshal(resMes)

	message := message.Message{
		Type: message.LoginResMesType,
		Data: string(resData),
	}

	data, err := json.Marshal(message)

	trans.WritePkg(data)

	return
}
