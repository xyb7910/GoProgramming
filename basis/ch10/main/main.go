package main

import (
	"LearingGo/basis/ch10/course"
	"fmt"
)

func main() {
	c := course.Course{
		Name: "ppp",
	}
	fmt.Println(course.GetCourse(c))
}
