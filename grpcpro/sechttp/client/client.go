package client

import (
	"net/rpc"
)

// HelloServiceName 使用 包路径.服务名 作为服务名
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, response *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, response *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, response)
}
