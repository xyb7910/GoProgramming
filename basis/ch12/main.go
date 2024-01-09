package main

import (
	"fmt"
	"time"
)

//主协程
func main() {
	//主死随从

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(10 * time.Second)
}
