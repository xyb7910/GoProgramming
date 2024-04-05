package main

import (
	"fmt"
)

func main() {
	// 声明一个map

	//scoreMap := make(map[string]int, 10)
	//scoreMap["zhangsan"] = 100
	//scoreMap["lisi"] = 200
	//scoreMap["wangwu"] = 300
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["zhangsan"])
	//fmt.Println("type of a:%T\n", scoreMap)

	// 在声明时添加对象

	//scores := map[string]int{
	//	"zhangsan": 100,
	//	"lisi":     200,
	//	"wangwu":   300,
	//}
	//fmt.Println(scores)

	//判断某个key是否存在

	//v, ok := scores["xiaoyan"]
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("key not exist")
	//}

	// 遍历map

	//for k, v := range scores {
	//	fmt.Println(k, v)
	//}

	// 删除map中的某个key

	//delete(scores, "zhangsan")
	//v, ok := scores["zhangsan"]
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("key not exist")
	//}

	// 元素为map的切片
	//mapSlice := make([]map[string]int, 3)
	//for k, v := range mapSlice {
	//	fmt.Println("index:", k, "value:", v)
	//}
	//
	//// initialize map
	//mapSlice[0] = make(map[string]int, 2)
	//mapSlice[0]["zhangsan"] = 100
	//mapSlice[0]["lisi"] = 200
	//mapSlice[0]["wangwu"] = 300
	//for k, v := range mapSlice {
	//	fmt.Println("index:", k, "value:", v)
	//}
	//length := len(mapSlice[0])
	//cap := cap(mapSlice)
	//fmt.Println(length)
	//fmt.Println(cap)

	// 值为slice的map
	//sliceMap := make(map[string][]int, 3)
	//fmt.Println(sliceMap)
	//sliceMap["zhangsan"] = []int{100, 200, 300}
	//sliceMap["zhangsan"] = []int{100, 200, 300, 400}
	//sliceMap["lisi"] = []int{100, 200, 300, 400, 500}
	//value, ok := sliceMap["zhangsan"]
	//if ok {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("key not exist")
	//}

	// 统计字符串中的字符出现次数
	//str := "how do you do"
	//countMap := make(map[string]int, 26)
	//strArr := strings.Split(str, "")
	//fmt.Println(strArr)
	//for _, v := range strArr {
	//	if v == " " {
	//		continue
	//	} else {
	//		if _, ok := countMap[v]; ok {
	//			countMap[v]++
	//		} else {
	//			countMap[v] = 1
	//		}
	//	}
	//}
	//fmt.Println(countMap)

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3) // 1 2 3
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...) // 1 3
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
