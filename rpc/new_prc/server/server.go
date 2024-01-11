package main

import (
	"LearingGo/rpc/new_prc/handler"
	"LearingGo/rpc/new_prc/serverproxy"
	"net"
	"net/rpc"
)

func main() {
	hellohandler := &handler.HelloServer{}
	_ = serverproxy.RegisterHelloService(hellohandler)
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("调用失败")
	}
	conn, err := lis.Accept()
	if err != nil {
		panic("建立连接")
	}
	rpc.ServeConn(conn)
}
