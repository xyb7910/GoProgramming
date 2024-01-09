package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//长度计算
	//如果您想要知道一个字符串（中文）长度，如果只有中文，可以使用len
	name := "Hello, World!"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转译符
	coursename := `go "体系课"` //ESC下边的键
	fmt.Println(coursename)

	//格式化输出
	username := "YPB"
	age := 18
	address := "北京"
	moblie := "18429033729"
	fmt.Println("用户名："+username, ", 年龄："+strconv.Itoa(age)+", 地址："+address, ", 电话："+moblie)
	fmt.Printf("用户名：%s, 年龄：%d, 地址：%s, 电话：%s\r\n", username, age, address, moblie)
	//记住
	useMsg := fmt.Sprintf("用户名：%s, 年龄：%d, 地址：%s, 电话：%s\r\n", username, age, address, moblie)
	fmt.Println(useMsg)

	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	builder.WriteString("年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("地址：")
	builder.WriteString(address)
	builder.WriteString("电话：")
	builder.WriteString(moblie)

	re := builder.String()
	fmt.Println(re)
}
