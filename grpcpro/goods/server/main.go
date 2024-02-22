package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct{}

// AddGoodsReq add goods request struct
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

// AddGoodsRes add goods response struct
type AddGoodsRes struct {
	Success bool
	Message string
}

// GetGoodsReq get goods request struct
type GetGoodsReq struct {
	Id int
}

// GetGoodsRes get goods response struct
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (g *Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	fmt.Printf("add goods: %v\n", req)
	*res = AddGoodsRes{
		Success: true,
		Message: "add success",
	}
	return nil
}

func (g *Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	fmt.Printf("get goods: %v\n", req)
	*res = GetGoodsRes{
		Id:      12,
		Title:   "服务器获取的数据",
		Price:   100,
		Content: "服务器获取的数据",
	}
	return nil
}

func main() {
	// 1.创建rpc server
	err := rpc.RegisterName("Goods", new(Goods))
	if err != nil {
		fmt.Printf("register rpc error: %v\n", err)
	}

	// 2.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("listen error: %v\n", err)
	}
	defer listener.Close()

	for {
		fmt.Println("wait for connection...")
		// 3.accept连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept error: %v\n", err)
			return
		}
		// 4.处理连接
		rpc.ServeConn(conn)
	}

}
