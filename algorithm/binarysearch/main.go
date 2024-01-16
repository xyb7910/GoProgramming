package main

import (
	"LearingGo/algorithm/binarysearch/src"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 5, 8, 9, 10}
	target := 5
	idx := src.BinarySearchMatrix(nums, target)
	fmt.Println("target is", nums[idx])
	firstSearchTarget := src.SearchFirstEqualElement(nums, target)
	fmt.Println("first equal element local is", firstSearchTarget)
	lastSearchTarget := src.SearchLastEqualElement(nums, target)
	fmt.Println("last equal element local is", lastSearchTarget)
	firstGreaterElement := src.SearchFirstGreaterElement(nums, target)
	fmt.Println("first greater element local is", firstGreaterElement)
	lastLastLessElement := src.SearchLastLessElement(nums, target)
	fmt.Println("last less element local is", lastLastLessElement)

}
