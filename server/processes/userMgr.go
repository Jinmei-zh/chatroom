package processes

import "fmt"

// 在线用户模块
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 添加在线用户
func (this *UserMgr) AddOnlineUser(user *UserProcess) {
	this.onlineUsers[user.UserId] = user
}

// 删除在线用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 查询所有用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 查询所有用户id列表
func (this *UserMgr) GetAllOnlineUserIds() (ids []int) {
	for id := range this.onlineUsers {
		ids = append(ids, id)
	}
	return
}

// 查询单个用户信息
func (this *UserMgr) GetOnlineUserById(userId int) (user *UserProcess, err error) {
	user, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不在线\n", userId)
		return
	}
	return
}
