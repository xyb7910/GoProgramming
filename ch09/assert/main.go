package main

import "fmt"

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int) //断言 对不同类型的代码
		bi, _ := b.(int) //断言
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("not suppprted type")
	}
}

func main() {
	a := "abc"
	b := "def"
	fmt.Println(add(a, b))
}
