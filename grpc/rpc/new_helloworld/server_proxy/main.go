package server_proxy

import (
	"LearingGo/grpc/rpc/new_helloworld/handler"
	"net/rpc"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
