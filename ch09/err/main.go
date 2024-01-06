package main

import (
	"fmt"
)

func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mPrint1(data interface{}) {
	fmt.Println(data)
}

func main() {
	var data = []interface{}{
		"yxc", 121, 1.80,
	}
	mPrint(data)

	var data1 = []string{
		"yxc", "ypb", "abc",
	} //不可以

	var data2 []interface{} //曲线救国
	for _, value := range data1 {
		data2 = append(data2, value)
	}
	fmt.Println(data2)

}
