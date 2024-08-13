package work

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "path/to/proto/hello"

func DoClientWork(client *rpc.Client) {
	helloCall := client.Go(HelloServiceName+".SayHello", "ypb", new(string), nil)
	// 处理其他业务
	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatalf("rpc call error: %v", err)
	}
	args := helloCall.Args.(*string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, reply)
}
