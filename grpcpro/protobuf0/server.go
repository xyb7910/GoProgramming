package protobuf0

import (
	"fmt"
	pro "grpcpro/protobuf0/proto"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/proto/hello"

type HelloServer struct{}

type HelloServiceInterface interface {
	SayHello(request *pro.String, response *pro.String) error
}

func (p *HelloServer) SayHello(request *pro.String, response *pro.String) error {
	response.Value = "hello" + request.GetValue()
	return nil
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func SimpleRpcServer() {
	_ = RegisterHelloService(new(HelloServer))

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("net.Listen error:", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
