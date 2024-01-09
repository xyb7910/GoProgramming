package main

import "fmt"

type Person1 struct {
	name    string
	age     int
	address string
	height  float32
}
type Person struct {
	name string
	age  int
}

type Student struct {
	p     Person
	score float32
}

func (s Student) print() {
	fmt.Printf("name: %s, age: %d", s.p.name, s.p.age)
}

func main() {
	/*
		p1 := Person{"yxc", 18, "Xian", 180.2}
		fmt.Println(p1)

		p2 := Person{
			name:    "ypb",
			age:     18,
			address: "Beijing",
			height:  190,
		}
		fmt.Println(p2)

		var persons []Person
		persons = append(persons, p1)
		persons = append(persons, Person{
			name: "zxc",
		})
		fmt.Println(persons)
	*/

	/*
		var p Person
		p.name = "yyy"
		fmt.Println(p.name)
		fmt.Println(p.age)

		//匿名结构体
		address := struct {
			province string
			city     string
			addres   string
		}{
			"beijing",
			"tongzhouqu",
			"XXX",
		}
		fmt.Println(address)
	*/

	//结构体的嵌套

	s := Student{
		Person{"yxc", 18},
		100,
	}
	fmt.Println(s)

	s1 := Student{}
	s1.p.name = "ypb"
	s1.p.age = 19
	s1.score = 98
	fmt.Println(s1)

	s1.print()
}
