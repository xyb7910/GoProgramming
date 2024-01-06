package main

import (
	u10 "LearingGo/ch10/course" //常用别名机制,也可以使用匿名机制来实现导入但不使用
	"fmt"
)

func main() {
	c := u10.Course{
		Name: "ppp",
	}
	fmt.Println(u10.GetCourse(c))
}
