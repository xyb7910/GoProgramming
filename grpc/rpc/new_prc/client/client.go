package main

import (
	"LearingGo/grpc/rpc/new_prc/clientproxy"
	"fmt"
)

func main() {
	client := clientproxy.NewClient("localhost:1234")
	var reply string
	err := client.Hello("ypb", &reply)
	if err != nil {
		panic("调用服务失败")
	}
	fmt.Println(reply)
}
