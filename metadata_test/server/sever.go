package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"LearingGo/metadata_test/proto"
)

type Greeter struct{}

func (g *Greeter) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("error")
	}
	for k, v := range md {
		fmt.Println(k, v)
	}
	if nameSlice, ok := md["name"]; ok {
		for _, val := range nameSlice {
			fmt.Println(val)
		}
	}
	if passwordSlice, ok := md["password"]; ok {
		for _, val := range passwordSlice {
			fmt.Println(val)
		}
	}
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}
func main() {
	g := grpc.NewServer()
	s := Greeter{}
	proto.RegisterGreeterServer(g, &s)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(lis)
}
