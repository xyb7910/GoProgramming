package client

import (
	"fmt"
	"log"
	"testing"
)

func TestDialHelloService(t *testing.T) {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
