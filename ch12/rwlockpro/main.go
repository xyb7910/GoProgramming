package main

import (
	"fmt"
	"sync"
)

var (
	mu   sync.RWMutex
	data = make(map[int]int)
	wg   sync.WaitGroup
)

func writeData(i int) {
	mu.Lock()
	defer mu.Unlock()
	data[i] = i
}

func readData(i int) {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("读取数据：", data[i])
}

func main() {
	wg.Add(10)

	// 创建写Goroutine
	go func() {
		for i := 0; i < 5; i++ {
			writeData(i)
		}
		wg.Done()
	}()

	// 创建读Goroutine
	for i := 0; i < 5; i++ {
		go func(i int) {
			readData(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
