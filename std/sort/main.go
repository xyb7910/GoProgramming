package main

import (
	"fmt"
	"sort"
)

type NewIntSlice []uint

func (s NewIntSlice) Len() int {
	return len(s)
}

func (s NewIntSlice) Less(i, j int) bool {
	fmt.Println(i, j, s[i] < s[j], s)
	return s[i] < s[j]
}

func (s NewIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := []uint{4, 2, 10, 1, 5, 8, 7, 6, 9, 3}
	sort.Sort(NewIntSlice(s))
	fmt.Println(s)
}
