package gorpc

import (
	"fmt"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: client}, nil
}

func (h *HelloServiceClient) SayHello(req string, res *string) error {
	return h.Client.Call(HelloServiceName+".SayHello", req, res)
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func SimpleClientV2() {
	client, err := DialHelloService("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Dial err:", err)
	}
	var reply string
	err = client.SayHello("ypb", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
	}
	fmt.Println("reply:", reply)
}

func SimpleClientV1() {
	// 1. 使用 gorpc 连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Dail err:", err)
		return
	}
	defer conn.Close()

	// 2. 调用服务器的接口
	var reply string
	// Call 第一个参数为服务器接口名，第二个参数为参数，最后一个参数为返回值
	err = conn.Call(HelloServiceName+".SayHello", "ypb", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}
	fmt.Println("reply:", reply)
}
