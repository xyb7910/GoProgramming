// 代码运行时出现invalid receiver type ClientCodec (ClientCodec is an interface type)，帮我修改一下
package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改replay的值
	*reply = "hello," + request
	return nil
}

func main() {
	//1、实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2、注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3、启动服务
	conn, _ := listener.Accept() //当一个新的连接进来的时候
	rpc.ServeConn(conn)
}
