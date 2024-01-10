package main

import (
	"LearingGo/rpc/new_helloworld/handler"
	"LearingGo/rpc/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	//1、实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2、注册处理逻辑 handler
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	//3、启动服务
	for {
		conn, _ := listener.Accept() //当一个新的连接进来的时候
		go rpc.ServeConn(conn)
	}

}
