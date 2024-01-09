package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8 = 12
	var b = uint8(a)
	fmt.Println(b)

	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(f, c)

	type IT int //类型要求很严格
	var it = IT(78)
	fmt.Println(it)

	//字符串转数字
	var istr = "12"
	myint, err := strconv.Atoi(istr) //返回两个值
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	//数字转字符串
	var myi = 32
	mstr := strconv.Itoa(myi)
	fmt.Println(mstr)

	//字符串转换为float32， 转换为bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	parseint, err := strconv.ParseInt("12", 10, 64)
	if err != nil {
		return
	}
	fmt.Println(parseint)

	parabool, err := strconv.ParseBool("true1")
	if err != nil {
		fmt.Println("ParaBool error")
	}
	fmt.Println(parabool)

	//基本类型转字符串
	boolstr := strconv.FormatBool(true)
	fmt.Println(boolstr)

	floatstr := strconv.FormatFloat(3.1415926, 'f', -1, 64)
	fmt.Println(floatstr)

	intstr := strconv.FormatInt(32, 16)
	fmt.Println(intstr)
}
