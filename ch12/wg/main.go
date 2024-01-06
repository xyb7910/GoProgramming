package main

import (
	"fmt"
	"sync"
)

//子 goroutine 如何通知到主的 goroutine 自己结束了， 主的goroutine如何知道 子的goroutine 已经结束了

func main() {
	var wg sync.WaitGroup

	//我要监控多少个 goroutine 执行
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("all done")

	//waitgroup 主要用于 goroutine 的执行等待  Add方法要和Done方法配套
}
