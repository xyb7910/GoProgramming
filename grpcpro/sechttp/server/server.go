package main

import (
	"LearingGo/sechttp/client"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, response *string) error {
	*response = "Hello " + request
	return nil
}

func main() {
	client.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, er := listener.Accept()
		if er != nil {
			log.Fatal("Accept error:", er)
		}
		go rpc.ServeConn(conn)
	}
}
