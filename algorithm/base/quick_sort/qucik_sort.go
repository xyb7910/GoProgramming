package main

import "fmt"

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left-1, right+1
	x := nums[(left+right)/2]
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	QuickSort(nums, left, j)
	QuickSort(nums, j+1, right)
}

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	QuickSort(nums, 0, n-1)

	fmt.Println(nums[k-1])
	fmt.Println(" ")
}
