package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {
	// 1.使用 rpc client 调用 rpc server
	conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 2.调用rpc server的rpc方法
	var reply AddGoodsRes
	err = conn.Call("Goods.AddGoods", AddGoodsReq{
		Id:      1,
		Title:   "test",
		Price:   100,
		Content: "test",
	}, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("reply: %#v\n", reply)

	var reply2 GetGoodsRes
	err = conn.Call("Goods.GetGoods", GetGoodsReq{
		Id: 11,
	}, &reply2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("reply2: %#v\n", reply2)
}
