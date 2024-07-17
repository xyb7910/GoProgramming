package main

import "errors"

/*
实现切片的相关操作：
 1. 实现切片的删除操作
*/

var ErrIndexOutOfRange = errors.New("index out of range")

// DeleteSliceByIndex 删除切片中下标为index的元素
func DeleteSliceByIndex[T any](slice []T, index int) ([]T, error) {
	length := len(slice)

	// 进行边界检查
	if index < 0 || index >= length {
		return nil, ErrIndexOutOfRange
	}

	// 成功删除后，进行元素的移动
	for i := index; i+1 < length; i++ {
		slice[i] = slice[i+1]
	}
	// 返回结果，并且删除最后一个元素
	return slice[:length-1], nil
}

// 考虑切片缩容
