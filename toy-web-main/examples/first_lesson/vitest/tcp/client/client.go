package main

import (
	"fmt"
	"net"
)

func main() {

	//连接服务端的地址
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接服务端成功", conn.RemoteAddr())

}
