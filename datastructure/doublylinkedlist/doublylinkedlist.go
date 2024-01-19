package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) InsertFirst(v int) {
	data := &Node{data: v}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList) InsertLast(v int) {
	data := &Node{data: v}
	if list.head == nil {
		list.head = data
		list.tail = data
	}
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.tail
	}
	list.tail = data
}

func (list *LinkedList) RemoveByValue(v int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == v {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == v {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == v {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return false
	}
	current := list.head
	for u := 0; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList) SearchValue(v int) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == v {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

func (list *LinkedList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

func (list *LinkedList) GetSize() int {
	size := 0
	current := list.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (list *LinkedList) GetItemsFromStart() []int {
	var items []int
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *LinkedList) GetItemsFromEnd() []int {
	var items []int
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}

func main() {
	dlist := LinkedList{}

	for i := 1; i <= 5; i++ {
		dlist.InsertFirst(i)
	}
	fmt.Println(dlist.GetSize())
	fmt.Println(dlist.GetItemsFromStart())
	fmt.Println(dlist.GetItemsFromEnd())

	dlist.RemoveByValue(3)
	dlist.RemoveByIndex(0)
	fmt.Println(dlist.GetSize())
	fmt.Println(dlist.SearchValue(3))

}
