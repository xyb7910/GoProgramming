package main

import "fmt"

type Slice[T int | string] []T
type Map[K int | string, V int | string] map[K]V

func reverse(s []int) []int {
	l := len(s)
	r := make([]int, l)
	// 新建一个数组用于存储反转后的结果
	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := reverseWithGenerics([]int{1, 2, 3, 4, 5})
	for _, e := range r {
		fmt.Println(e)
	}
	fmt.Println(min[int](1, 2))
	fmt.Println(min[float64](1.0, 2.0))
}
