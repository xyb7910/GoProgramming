package main

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (list *Stack) Push(v int) {
	data := &Node{data: v}
	if list.top != nil {
		data.next = list.top
	}
	list.top = data
}
