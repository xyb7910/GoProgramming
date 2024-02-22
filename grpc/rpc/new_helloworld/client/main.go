package main

import (
	"LearingGo/grpc/rpc/new_helloworld/client_proxy"
	"fmt"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	//1、建立连接
	var reply string
	err := client.Hello("ypb", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
