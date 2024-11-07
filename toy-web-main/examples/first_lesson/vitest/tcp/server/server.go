package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("服务器开始监听")
	listener, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println("listen err=", err)
	}
	defer listener.Close()

	//循环连接等到客户端来连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err=", err) //其中一个连接失败了　但是不需要退出
		} else {
			fmt.Printf("accept() succes con=%v,客户端的ip＝\n", conn, conn.RemoteAddr().String())
		}
		//这里要准备起来一个协程为客户端服务

	}

	fmt.Println("listen succes=", listener)
}
