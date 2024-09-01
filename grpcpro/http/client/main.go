package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "ypb", &reply)
	if err != nil {
		log.Fatal("calling:", err)
	}
	fmt.Println(reply)
}
