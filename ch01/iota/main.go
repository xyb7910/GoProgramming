package main

import "fmt"

func a() (int, bool) {
	return 0, false
}

func main() {
	//匿名变量
	var _ int
	_, ok := a()
	if ok {
		//打印
	}
	const (
		ERR1 = iota + 1
		ERR2 // iota内部计数器会自动计数
		ERR3 = "S"
		ERR4
	)
	fmt.Println(ERR1)
	fmt.Println(ERR2)
	fmt.Println(ERR3)
	fmt.Println(ERR4)

	/*
		如果中断了iota的定义必须显式的恢复，后续会自动递增
		自增类型默认是int类型
		iota能弱化const类型的定义
	*/
}
