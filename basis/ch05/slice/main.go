package main

import "fmt"

func main() {
	/*
		// 折中 slice 切片
		var courses []string
		fmt.Printf("%T\n", courses)

		//使用append追加
		courses = append(courses, "go")
		courses = append(courses, "grpc")
		courses = append(courses, "gin")

		fmt.Println(courses[1])

		// 切片初始化三种方式：1.从数组直接创建 2.使用string{} 3. make

		allCourses := [5]string{"go", "grpc", "mysql", "c++"}
		courseSlice := allCourses[0:len(allCourses)]
		fmt.Println(courseSlice)


	*/
	courseSlice := []string{"go", "grpc", "mysql", "c++"}
	/*
		//使用make函数
		allCourses2 := make([]string, 3)
		allCourses2[0] = "c"
		allCourses2[1] = "c"
		allCourses2[2] = "c"
		//allCourses2[3] = "c"
		fmt.Println(allCourses2)

		var allCourses3 []string
		allCourses3 = append(allCourses3, "c");
		fmt.Println(allCourses3)
	*/

	/*
		//访问切片元素，访问单个，访问多个
		fmt.Println(courseSlice[1])
		fmt.Println(courseSlice[0:2]) //左开右闭
		fmt.Println(courseSlice[0:])  //只有start没有end， 表示从start到结尾所有元素
		fmt.Println(courseSlice[:3])  //只有end没有start，表示从0到end所有元素
		fmt.Println(courseSlice[:])   //没有start也没有end，表示全部输出

		//将两个slice进行追加，注意在随后一个后边添加...
		courseSlice = append(courseSlice, "c", "python")
		courseSlice1 := []string{"ruby", "rust"}
		courseSlice = append(courseSlice, courseSlice1[:]...)
		fmt.Println(courseSlice)
	*/

	// 删除slice中的元素，重新生成
	myslice := append(courseSlice[0:1], courseSlice[2:]...) //删除grpc
	fmt.Println(myslice)
	courseSlice = courseSlice[0:3]
	fmt.Println(courseSlice)

	//拷贝slice
	courseSliceCopy := courseSlice
	courseSliceCopy1 := courseSlice[:]
	fmt.Println(courseSliceCopy)
	fmt.Println(courseSliceCopy1)

	var courseSliceCopy2 = make([]string, len(courseSlice))
	copy(courseSliceCopy2, courseSlice)
	fmt.Println(courseSliceCopy2)
	
	courseSlice[0] = "java"
	fmt.Println(courseSliceCopy2)
	fmt.Println(courseSliceCopy1)
}
