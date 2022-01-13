package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) WriteReply(data []byte) (mes message.Message, err error) {
	mes, err = this.ReadPkg()
	if err != nil {
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// fmt.Println("等待读取接收数据...")
	conn := this.Conn
	_, err = conn.Read(this.Buf[:4])
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("读取长度数据错误 err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	_, err = conn.Read(this.Buf[:pkgLen])
	if err != nil {
		fmt.Println("读取主数据错误 err=", err)
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("数据解析错误", err)
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 发送长度
	pkgLen := uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	nLen, err := this.Conn.Write(this.Buf[0:4])
	if nLen != 4 || err != nil {
		fmt.Println("长度发送错误  err=", err)
		return
	}

	// 发送消息
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("发送数据错误  err=", err)
		return
	}
	return
}
