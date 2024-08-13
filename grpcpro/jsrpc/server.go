package jsrpc

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "path/to/proto/hello"

type HelloServer struct{}

type HelloServerInterface interface {
	SayHello(request string, resp *string) error
}

func (h *HelloServer) SayHello(request string, resp *string) error {
	*resp = "Hello " + request
	return nil
}

func RegisterHelloServer(svc HelloServerInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func SimpleServerV2() {
	_ = rpc.Register(new(HelloServer))
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	_ = http.ListenAndServe(":1234", nil)
}

func SimpleServerV1() {
	_ = RegisterHelloServer(new(HelloServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
