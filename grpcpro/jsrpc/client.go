package jsrpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func SimpleClient() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		fmt.Printf("dial err: %v\n", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(HelloServiceName+".SayHello", "ypb", &reply)
	if err != nil {
		fmt.Printf("call err: %v\n", err)
	}
	fmt.Printf("reply: %v\n", reply)
}
