package message

type User struct {
	Id     int    `json:"id"`
	Pwd    string `json:"pwd"`
	Name   string `json:"name"`
	Status int    `json:"status"` //1在线 2离线
}
