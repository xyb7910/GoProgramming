package main

import "fmt"

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from"+"goroutine %d\n",
					i)
			}
		}(i)
	}
}
