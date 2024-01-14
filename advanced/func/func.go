package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unopported opteration: %s", op)

	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b //返回除数，余数
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += i
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()    //获取op函数所在内存中的地址
	opName := runtime.FuncForPC(p).Name() //获取op函数具体名称
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func main() {
	if result, err := eval(2, 3, "+"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))

	fmt.Printf("1+2+3...+5", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}
