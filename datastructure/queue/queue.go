package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	rear *Node
}

// 入队
func (list *Queue) Enqueue(v int) {
	data := &Node{data: v}
	if list.rear != nil {
		data.next = list.rear
	}
	list.rear = data
}

// 出队
func (list *Queue) Dequeue() (int, bool) {
	if list.rear == nil {
		return 0, false
	}
	if list.rear.next == nil {
		i := list.rear.data
		list.rear = nil
		return i, true
	}

	current := list.rear
	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, false
		}
		current = current.next
	}
}

func (list *Queue) Peek() (int, bool) {
	if list.rear == nil {
		return 0, false
	}
	return list.rear.data, true
}

func (list *Queue) Get() []int {
	var items []int
	current := list.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *Queue) IsEmpty() bool {
	return list.rear == nil
}

func (list *Queue) Empty() {
	list.rear = nil
}

func main() {
	q := Queue{}

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	res := q.Get()
	fmt.Println(res)

	fmt.Println(q.IsEmpty())
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	fmt.Println(q.Get())
	q.Empty()
	fmt.Println(q.IsEmpty())

}
