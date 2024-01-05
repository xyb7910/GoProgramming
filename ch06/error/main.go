package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	return 0, errors.New("This is a error!")
}

func main() {
	//painc 会导致整个程序退出，不推荐使用painc

	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}
