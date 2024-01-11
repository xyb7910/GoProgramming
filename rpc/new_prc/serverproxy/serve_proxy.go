package serverproxy

import "net/rpc"

const HelloServerName = "handler/HelloServer"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServerName, srv)
}
