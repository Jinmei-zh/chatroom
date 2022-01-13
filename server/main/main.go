package main

import (
	"chatroom/server/model"
	"chatroom/server/processes"
	"chatroom/server/utils"
	"fmt"
	"net"
	"time"

	"github.com/garyburd/redigo/redis"
)

func process(conn net.Conn) {
	defer conn.Close()

	process01 := &processes.Processor{
		Conn: conn,
	}
	err := process01.MainProcess()
	if err != nil {
		fmt.Println("退出协程 error=", err)
		return
	}

	return
}

func initUserDao(pool *redis.Pool) {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 初始化redis连接池
	pool := utils.InitPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initUserDao(pool)

	// chatroom.TestNets()
	const Address = "0.0.0.0:8888"
	fmt.Println("服务器开启，地址：", Address)

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println("lister err=", err)
	}
	defer listen.Close()
	for {
		fmt.Println("等待连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err=", err)
		} else {
			fmt.Println("连接地址：", conn.LocalAddr().String())
		}

		go process(conn)
	}

}
