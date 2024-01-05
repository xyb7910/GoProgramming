package main

import "fmt"

func main() {
	// map 是一个key：value的无序集合,主要是查询方便  	必须进行初始化，才能使用
	//两种初始化
	var courseMap1 = map[string]string{}        //初始化一个空的map
	var courseMap2 = make(map[string]string, 3) // make是一个内置函数，主要用于初始化slice，map， channel
	fmt.Println(courseMap1)
	fmt.Println(courseMap2)

	var m []string
	if m == nil { //nil 与其他语言的null不同
		fmt.Println("yes")
	}
	m = append(m, "a")

	var courseMap = map[string]string{
		"go":  "yxc",
		"c++": "ypb",
		"gin": "gtq",
	}
	//取值
	fmt.Println(courseMap["go"])

	//放值
	courseMap["python"] = "ppp"
	fmt.Println(courseMap["python"])

	//遍历map ，每次编译打印是顺序不确定
	for _, value := range courseMap {
		fmt.Println(value)
	}

	for key := range courseMap {
		fmt.Println(key, courseMap[key])
	}

	//判断元素是否存在
	data, ok := courseMap["python"] //注意两个参数
	if !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("find ", data)
	}

	if _, ok := courseMap["java"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("in")
	}

	//删除一个元素
	delete(courseMap, "python")
	delete(courseMap, "c") //删除不存在的元素并不会报错

	//很重要的提示，map的线程并不安全
	fmt.Println(courseMap)
}
