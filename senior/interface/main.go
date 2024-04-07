package main

import "fmt"

//// print interface
//func print(i interface{}) {
//	fmt.Println(i)
//}
//
//func main() {
//	var a interface{}
//	a = 2
//	fmt.Printf("%T, %v\n", a, a)
//
//	// inject interface
//	print(a)
//	print("hello")
//	print(true)
//
//	// use assert to check int type
//	v, ok := a.(int)
//	if ok {
//		fmt.Printf("a is a int type, value is %d\n", v)
//	}
//
//	// use assert to check type
//	switch a.(type) {
//	case int:
//		fmt.Println("a is a int type")
//	case string:
//		fmt.Println("a is a string type")
//	case bool:
//		fmt.Println("a is a bool type")
//	default:
//		fmt.Println("a is a other type")
//	}
//
//	// use reflect to find type
//	t := reflect.TypeOf(a)
//	fmt.Printf("a is type: %s", t.Name())
//
//}

var (
	x Mover = new(Dog)
	y Mover = new(Car)
)

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("狗在跑~")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑~")
}

func main() {
	var m Mover
	dog := &Dog{Name: "旺财"}
	dog.Move()
	fmt.Println(m == nil)
	fmt.Println(x == y)
}
