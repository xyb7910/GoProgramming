package clientproxy

import "net/rpc"

const HelloServerName = "handler/HelloServer"

type HelloServiceClient struct {
	*rpc.Client
}

func NewClient(address string) HelloServiceClient {
	conn, err := rpc.Dial("tcp", address)
	if err != nil {
		panic("连接服务器失败")
	}
	return HelloServiceClient{conn}
}

func (c *HelloServiceClient) Hello(request string, reply *string) error {
	err := c.Call(HelloServerName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil

}
