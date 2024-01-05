package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changname(p *Person) {
	p.name = "yxc"
	fmt.Printf("%p\n", p)
}

func swap(a, b *int) (*int, *int) {
	t := *a
	*a = *b
	*b = t
	return a, b
}

func main() {
	/*
		//指针， 结构体在传值的时候，在函数中修改的值可以反应到变量中 , go	的指针不能进行 + 1
		p := Person{
			name: "ypb",
			age:  18,
		}
		fmt.Println(p.name)
		changname(&p)
		fmt.Println(p.name)

		a := 10
		b := &a
		fmt.Println(a, b)

		ps := &Person{} //第一种初始化 ， 指针必须进行初始化
		fmt.Println(ps)
		var emptyPerson *Person // 第二种初始化 ，
		pi := &emptyPerson
		fmt.Println(pi)
		var pp = new(Person) //第三种，指针初始化推荐方法
		fmt.Println(pp.name)
	*/

	a := 10
	b := 5
	swap(&a, &b)
	fmt.Println(a, b)
}
