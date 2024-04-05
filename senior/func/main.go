package main

import (
	"errors"
	"fmt"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "add":
		return add, nil
	case "sub":
		return sub, nil
	case "mul":
		return mul, nil
	default:
		err := errors.New("invalid operation")
		return nil, err
	}
}

// 函数的闭包
func makeSuffixFunc(suffix string) func(string) string {
	// 返回一个函数 -- 匿名函数
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//func main() {
//	do, err := do("add")
//	a := do(1, 2)
//	if err == nil {
//		fmt.Println(a)
//	}
//
//	jpgFunc := makeSuffixFunc(".jpg")
//	txtFunc := makeSuffixFunc(".txt")
//	fmt.Println(jpgFunc("test"))
//	fmt.Println(txtFunc("test"))
//}

//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++
//	}()
//	return 5
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	defer func(x int) {
//		x++
//	}(x)
//	return 5
//}
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}

// defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//
//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//}

//func funcA() {
//	fmt.Println("func A")
//}
//
//func funcB() {
//	defer func() {
//		err := recover()
//		//如果程序出出现了panic错误,可以通过recover恢复过来
//		if err != nil {
//			fmt.Println("recover in B")
//		}
//	}()
//	panic("panic in B")
//}
//
//func funcC() {
//	fmt.Println("func C")
//}
//func main() {
//	funcA()
//	funcB()
//	funcC()
//}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,

	Peter,Giana,Adriano,Aaron,Elizabeth。

分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coinsSum = 50
	users    = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatch(name string) (res int) {
	if strings.Contains(name, "e") || strings.Contains(name, "E") {
		res = 1
	} else if strings.Contains(name, "i") || strings.Contains(name, "I") {
		res = 2
	} else if strings.Contains(name, "o") || strings.Contains(name, "O") {
		res = 3
	} else if strings.Contains(name, "u") || strings.Contains(name, "U") {
		res = 4
	}
	return
}

// distribution 返回分配之后剩余的金币数
func dispatchCoin() int {
	for _, user := range users {
		if _, ok := distribution[user]; ok {
			distribution[user] += dispatch(user)
		} else {
			distribution[user] = dispatch(user)
		}

	}
	var sum int
	for _, v := range distribution {
		sum += v
	}
	return coinsSum - sum
}

func dispatchCoinAI(names []string) (map[string]int, int) {
	for _, name := range names {
		coins := 0
		for _, char := range strings.ToLower(name) {
			switch char {
			case 'e':
				coins += 1
			case 'i':
				coins += 2
			case 'o':
				coins += 3
			case 'u':
				coins += 4
			}
		}
		distribution[name] = coins
		coinsSum -= coins
	}

	return distribution, coinsSum
}

func main() {
	_, left := dispatchCoinAI(users)
	for k, v := range distribution {
		fmt.Println(k, v)
	}
	fmt.Println("剩下：", left)
}
