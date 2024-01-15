package main

import (
	"LearingGo/algorithm/binarysearch/src"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	idx := src.BinarySearchMatrix(nums, target)
	fmt.Println("target is", nums[idx])
}
