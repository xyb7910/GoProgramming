package main

import "fmt"

type Person struct {
	name string
	age  int
}

var p1 = Person{
	name: "ypb",
	age:  19,
}
var p2 = Person{
	name: "ypb",
	age:  18,
}

func main() {
	/*
			不同类型的数据零值不一样
			bool false
			number 0
			string ""
			pointer nil
			slice nil
			map nil
			channel, interface, function nil

		struct 默认值不是nil，默认值是具体字段的默认值
	*/
	if p1 == p2 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	//slice的默认值
	var sli1 []Person
	//var ps2 = make([]Person, 0)
	if sli1 == nil {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}

	//map的默认值
	//var m map[string]string
	var m2 = make(map[string]string, 0)
	if m2 == nil {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}

	for _, value := range m2 {
		fmt.Println(value)
	}
}
