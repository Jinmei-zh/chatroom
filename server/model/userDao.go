package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) createUser(conn redis.Conn, id int, pwd string, name string) (err error) {
	// 先判断id是否存在
	_, err = this.getUserById(conn, id)
	if err != nil {
		// 写入新用户信息
		_, err = redis.String(conn.Do("HSET", "users", id))
		if err != nil {
			fmt.Println("写入数据失败 error=", err)
			return
		}
		return
	} else {
		err = ERROR_USER_EXISTS
		fmt.Println("已有id：", id)
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user message.User, err error) {
	// value, err := redis.Values(conn.Do("HGETALL", id))
	value, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	fmt.Printf("type=%T\n", user)

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user message.User, err error) {

	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.Pwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 新增
func (this *UserDao) Register(userId int, userPwd string, userName string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, userId)
	if err == nil { //用户已存在
		err = ERROR_USER_EXISTS
		return
	}

	err = this.createUser(conn, userId, userPwd, userName)
	if err != nil {
		return
	}
	return
}

// // 修改
// func (*User) Update() (err error) {

// 	return
// }

// // 删除
// func (*User) Delete() (err error) {

// 	return
// }
