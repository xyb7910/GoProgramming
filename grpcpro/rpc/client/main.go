package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 使用 rpc 连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Dail err:", err)
		return
	}
	defer conn.Close()

	// 2. 调用服务器的接口
	var reply string
	// Call 第一个参数为服务器接口名，第二个参数为参数，最后一个参数为返回值
	err = conn.Call("hello.SayHello", "ypb", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}
	fmt.Println("reply:", reply)
}
