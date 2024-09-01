package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, response *string) error {
	*response = "Hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}

	rpc.ServeConn(conn)
}
