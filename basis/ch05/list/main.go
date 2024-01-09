package main

import (
	"container/list"
	"fmt"
)

func main() {
	var mylist list.List

	//尾部放入元素
	mylist.PushBack("go")
	mylist.PushBack(343)
	mylist.PushBack("eye")

	//头部放入元素
	mylist.PushFront("c")

	//在位置 c 插入元素
	i := mylist.Front()
	for ; i != nil; i.Next() {
		if i.Value.(string) == "c" {
			break
		}
	}
	mylist.InsertBefore("hahah", i)
	mylist.Remove(i)

	//正序遍历
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//反向遍历
	for i := mylist.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
