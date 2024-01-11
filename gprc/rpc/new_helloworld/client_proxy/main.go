package client_proxy

import (
	"LearingGo/gprc/rpc/new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protcol, addres string) HelloServiceStub {
	conn, err := rpc.Dial(protcol, addres)
	if err != nil {
		panic("connect error")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return nil
	}
	return nil
}
