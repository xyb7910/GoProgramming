package main

import (
	"LearingGo/rpc/helloworld/proto"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "ypb",
		Age:     18,
		Courses: []string{"go", "gin", "hello"},
	}
	jsonStruct := Hello{Name: "ypb", Age: 18, Courses: []string{"go", "gin", "hello"}}
	jsonRsq, _ := json.Marshal(jsonStruct)
	fmt.Println(len(jsonRsq))
	rsq, _ := proto.Marshal(&req)
	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsq, &newReq)
	//fmt.Println(len(rsq))
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
