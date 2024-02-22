package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Hello 定义类对象
type Hello struct{}

// SayHello 实现接口并且绑定方法
func (h *Hello) SayHello(req string, res *string) error {
	*res = req + "你好！"
	return nil
}

/*
RPC 原则：
1、方法只能有两个可序列化的参数，其中第二个参数是指针类型,
参数的类型不能是channel（通道）、complex（复数类型）、func（函数）,
因为它们不能进行序列化

2、方法的返回值必须是error类型，如果没有返回值，则返回nil， 同时必须是公开的方法
*/

func main() {
	// 1. 创建rpc服务
	err := rpc.RegisterName("hello", &Hello{})
	if err != nil {
		fmt.Println("rpc.RegisterName error:", err)
		return
	}
	//2.设置监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("server start listen 1234 port")
		//3. 建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}
		// 4. 处理连接，绑定服务
		go rpc.ServeConn(conn)
	}
}
