package main

import (
	"fmt"
	"time"
)

func Hu() {
	time.Sleep(2 * time.Second)
	fmt.Printf("after 2 second hu!!\n")
}

func main() {
	go Hu()

	fmt.Printf("start hu!!\n")

	time.Sleep(3 * time.Second)
}
