package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁
//读进程之间应该并发，读和写之间应该串行，读和读之间也不应该并行

func main() {
	//var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)
	//写的 goroutine
	go func() {
		time.Sleep(time.Second * 3)
		defer wg.Done()
		rwlock.Lock() //加写锁，写锁会防止别的写锁的获取，和读锁的获取
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		time.Sleep(time.Second * 5)
		// num = 12
	}()

	//time.Sleep(time.Second)

	//读的 goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()

			for {
				rwlock.RLock() //加读锁，读锁不会阻止别人的读
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}
			//fmt.Println(num)
		}()
	}

	wg.Wait()
}
