package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
}

// 头部插入
func (list *Linkedlist) InsertFirst(v int) {
	data := &Node{data: v}
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

// 尾部插入
func (list *Linkedlist) InsertLast(v int) {
	data := &Node{data: v}
	if list.head == nil {
		list.head = data
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

// 按值删除(第一个)
func (list *Linkedlist) RemoveByValue(v int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == v {
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == v {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *Linkedlist) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head = list.head.next
		return true
	}

	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	return false
}

// 按值查找
func (list *Linkedlist) SearchValue(v int) bool {
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

// 获取首元素
func (list *Linkedlist) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

// 获取尾元素
func (list *Linkedlist) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

// 获取长度
func (list *Linkedlist) GetSize() int {
	size := 0
	current := list.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

// 获取所有元素
func (list *Linkedlist) GetItems() []int {
	var items []int
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func main() {
	list := Linkedlist{}

	for i := 1; i <= 5; i++ {
		list.InsertFirst(i)
	}
	for i := 1; i <= 5; i++ {
		list.InsertLast(i)
	}

	fmt.Println(list.GetItems())

	fmt.Println(list.GetSize())
	list.RemoveByValue(5)
	fmt.Println(list.GetSize())

	fmt.Println(list.GetItems())
	list.RemoveByIndex(0)
	list.RemoveByIndex(list.GetSize() - 1)
	fmt.Println(list.GetItems())

	fmt.Println(list.SearchValue(3))
	fmt.Println(list.GetFirst())
	fmt.Println(list.GetLast())
}
