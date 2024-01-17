package main

import (
	"LearingGo/advanced/queue"
	"fmt"
)

func main() {
	q := queue.Queue{}

	q.Push(1)
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
