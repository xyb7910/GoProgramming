package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 整型切片
	//digitSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 字符切片
	charSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	index := 2
	back, err := DeleteSliceByIndex(charSlice, index)
	if err != nil {
		fmt.Println(ErrIndexOutOfRange)
	}
	for k, value := range back {
		fmt.Println(k, value)
	}

}
