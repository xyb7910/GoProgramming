package main

import (
	"fmt"
	"reflect"
)

/*
type MyInt int

type User struct {
	Name string
	Age  int
	Sex  string
}
*/

func main() {
	//var i int = 10
	//var s string = "hello"
	//var b bool = true
	//var f float64 = 1.23
	//var p *int = &i
	//var m map[string]int = map[string]int{"a": 1, "b": 2}

	/*
		fmt.Println(reflect.TypeOf(i))
		fmt.Println(reflect.TypeOf(s))
		fmt.Println(reflect.TypeOf(b))
		fmt.Println(reflect.TypeOf(f))
		fmt.Println(reflect.TypeOf(p))
		fmt.Println(reflect.TypeOf(m))
	*/

	/*
		reflectPrintType(i)
		reflectPrintType(s)
		reflectPrintType(b)
		reflectPrintType(f)
		reflectPrintType(p)
		reflectPrintType(m)
	*/

	/*
		var n MyInt = 10
		reflectPrintType(n) //TypeOf: main.MyInt  Name: MyInt Kind: int
		var u User = User{"Tom", 18, "male"}
		reflectPrintType(u) //TypeOf: main.User  Name: User Kind: struct
		var num = [3]int{1, 2, 3}
		reflectPrintType(num) //TypeOf: [3]int  Name:  Kind: array
	*/

	/*
		var a int64 = 10
		fmt.Printf("add befor: %v\n", a)
		reflectValue(a)
		fmt.Printf("add after: %v\n", a)
	*/

	/*
		var a int64 = 10
		reflectValue2(a)
		var s string = "hello"
		reflectValue2(s)
		var b struct {
			Name string
			Age  int
		}
		reflectValue2(b)
	*/

	var a int64 = 10
	reflectSetValue(a)
	fmt.Println(a)
	var b *int64 = &a
	reflectSetValue2(b)
	fmt.Println(a)

	/*
		x := 2
		v := reflect.ValueOf(&x)
		if v.CanSet() {
			v.Elem().SetInt(10)
		} else {
			fmt.Println("can not set")
		}
		fmt.Println(x)
	*/
}

func reflectSetValue(v interface{}) {
	fmt.Println(v)
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Int {
		value.SetInt(value.Int() + 6)
	}
	fmt.Println(v)
}

func reflectSetValue2(v interface{}) {
	value := reflect.ValueOf(v)

	// 在反射中通过 Elem（）来获取指针指向的值
	if value.Elem().Kind() == reflect.Int {
		value.Elem().SetInt(value.Elem().Int() + 6)
	}
}

func reflectValue2(v interface{}) {
	value := reflect.ValueOf(v)
	k := value.Kind()
	switch k {
	case reflect.Int64:
		// value.Int()
		fmt.Println("type is int64", value.Int())
	case reflect.Float64:
		fmt.Println("type is float64", value.Float())
	case reflect.String:
		fmt.Println("type is string", value.String())
	case reflect.Bool:
		fmt.Println("type is bool", value.Bool())
	default:
		fmt.Println("type is illegal")
	}
}

func reflectValue1(v interface{}) {
	value := reflect.ValueOf(v)
	var add = value.Int() + 6
	fmt.Println(add)
}

func reflectPrintType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("TypeOf: %v  Name: %v Kind: %v\n", t, t.Name(), t.Kind())
}
