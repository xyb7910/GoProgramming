package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
func main() {
	rpc.RegisterName("HelloServer", new(HelloServer))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("启动错误")
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic("接受")
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
