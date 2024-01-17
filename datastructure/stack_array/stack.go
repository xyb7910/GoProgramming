package main

import "fmt"

const MAXSIZE = 10

type Stack struct {
	top  int
	data [MAXSIZE]int
}

// 入栈
func (s *Stack) Push(v int) bool {
	if s.top == len(s.data) {
		return false
	}
	s.data[s.top] = v
	s.top++
	return true
}

// 出栈
func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	v := s.data[s.top-1]
	s.top--
	return v, true
}

// 获取栈顶元素
func (s *Stack) Peek() int {
	return s.data[s.top-1]
}

// 获取栈内所有元素
func (s *Stack) Get() []int {
	return s.data[:s.top]
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

// 栈置空
func (s *Stack) Empty() {
	s.top = 0
}

func main() {
	data := [MAXSIZE]int{}
	s := Stack{0, data}

	for i := 1; i <= 9; i++ {
		s.Push(i)
	}

	fmt.Println(s.Get())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Get())
	fmt.Println(s.Peek())
	fmt.Println(s.IsEmpty())
	s.Empty()
	fmt.Println(s.IsEmpty())

}
