package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 全局等待组

/*
var wg sync.WaitGroup

func main() {
	// 1. 创建一个等待组
	wg.Add(1)
	hello()
	fmt.Println("你好")
	// 2. 等待所有的协程执行完毕
	wg.Wait()
}

func hello() {
	fmt.Println("hello")
	// 3. 通知等待组已经执行完毕
	wg.Done()
}
*/

/*
var wg sync.WaitGroup

	func main() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go hello(i)
		}
		wg.Wait()
	}

	func hello(i int) {
		fmt.Println(i)
		wg.Done()
	}
*/

/*
var (
	wg   sync.WaitGroup
	x    int
	lock sync.Mutex
)

func Add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go Add()
	go Add()

	wg.Wait()
	fmt.Println(x)
}

*/

/*
// 测试读写锁与互斥锁的性能
var (
	x     int
	wg    sync.WaitGroup
	mutex sync.Mutex
	rwm   sync.RWMutex
)

func WriteWithMutex() {
	mutex.Lock()
	x++
	time.Sleep(time.Second)
	mutex.Unlock()
	wg.Done()
}

func ReadWithMutex() {
	mutex.Lock()
	time.Sleep(time.Second)
	mutex.Unlock()
	wg.Done()
}

func WriteWithRWMutex() {
	rwm.Lock()
	x = x + 1
	time.Sleep(time.Second)
	rwm.Unlock()
	wg.Done()
}

func ReadWithRWMutex() {
	rwm.RLock()
	time.Sleep(time.Second)
	rwm.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()

	// 写
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	// 读
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Println(cost)
}

func main() {
	do(WriteWithMutex, ReadWithMutex, 1, 6)
	do(WriteWithRWMutex, ReadWithRWMutex, 1, 6)
}
*/

// 测试sync.Map
var m = make(map[string]int) // map 不支持并发

func get(key string) int {
	res, ok := m[key]
	if ok {
		return res
	} else {
		fmt.Println("not found")
		return 0
	}
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	var smap = sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//set(key, n)
			smap.Store(key, n)
			//fmt.Println(get(key))
			value, _ := smap.Load(key)
			fmt.Printf("%v --- %v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
