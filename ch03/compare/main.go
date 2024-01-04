package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	b := "bello"

	//字符串比较
	fmt.Println(a == b)
	fmt.Println(a > b)

	name := "A-B-C-D-E-FG-A-B"
	//是否包含
	fmt.Println(strings.Contains(name, "AB"))
	//统计次数
	fmt.Println(strings.Count(name, "A"))
	//分割字符串
	fmt.Println(strings.Split(name, "-"))
	//字符串是否包含前缀
	fmt.Println(strings.HasPrefix(name, "A-B-C"))
	fmt.Println(strings.HasSuffix(name, "GR"))
	//查询子串出现的位置
	fmt.Println(strings.Index(name, "C"))
	//子串替换
	fmt.Println(strings.Replace(name, "-", "J", 2)) // -1 全部替换 1，2...分别表示替换几个
	//大小写转化
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))
	//去掉特殊字符
	fmt.Println(strings.Trim("  G O ", " "))      //修建两边
	fmt.Println(strings.TrimLeft("  G O ", " "))  //修建左边
	fmt.Println(strings.TrimRight("  G O ", " ")) //修建右边
}
