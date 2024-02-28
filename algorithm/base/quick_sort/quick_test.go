package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := 4
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums[n-1])
}
