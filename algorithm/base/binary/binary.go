package main

import "fmt"

func main() {
	var n, q, k int
	fmt.Scanf("%d%d", &n, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for ; q > 0; q-- {
		fmt.Scanf("%d", &k)
		l, r := 0, n-1
		for l < r {
			mid := l + r>>1
			if a[mid] >= k {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if a[l] != k {
			fmt.Printf("%d %d\n", -1, -1)
		} else {
			fmt.Printf("%d ", l)
			l, r = 0, n-1
			for l < r {
				mid := l + r + 1>>1
				if a[mid] <= k {
					l = mid
				} else {
					r = mid - 1
				}
			}
			fmt.Printf("%d\n", r)
		}
	}
}
