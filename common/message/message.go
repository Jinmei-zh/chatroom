package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "ResMessage"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SmsResMesType           = "SmsMes"
	SmsPrivateMesType       = "SmsPrivateMes"
	SignOutMesType          = "SignOutMes" //退出登陆
)
const (
	UserOnline  = 1 //在线
	UserOffline = 2 //离线
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	Id       int    `json:"id"`
	Password string `json:"pwd"`
}

type SignOutMes struct {
	UserId int `json:"user_id"`
}

type RegisterMes struct {
	User User `json:"data"`
}

type ResMessage struct {
	Code  int    `json:"code"`
	Error string `json:"message"`
}

type LoginResMes struct {
	Code     int    `json:"code"`
	Error    string `json:"message"`
	UserName string `json:"user_name"`
	Data     []int  `json:"data"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

// 群消息
type SmsMes struct {
	User
	Content string `json:"content"`
}

// 私信消息
type SmsPrivateMes struct {
	FromUser User
	Content  string `json:"content"`
	ToUserId int
}
