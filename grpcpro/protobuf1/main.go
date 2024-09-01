package main

import (
	pro "LearingGo/protobuf1/proto"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	user := &pro.Userinfo{
		Name:  "lzy",
		Age:   20,
		Hobby: []string{"basketball", "football"},
	}
	fmt.Println(user.GetHobby())

	// 使用 proto.Marshal() 将 Userinfo 转为 []byte
	data, err := proto.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	// 使用 proto.Unmarshal() 将 []byte 转为 Userinfo
	info := &pro.Userinfo{}
	err2 := proto.Unmarshal(data, info)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(info)
	fmt.Println(info.GetHobby())
}
