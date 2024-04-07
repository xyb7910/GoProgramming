package main

import (
	"fmt"
	"reflect"
)

type myInt int64

type person struct {
	name string
	age  int
}

type book struct {
	tittle string
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(100) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}
}

type student struct {
	Name string
	Age  int
}

func main() {

	//// *int 类型的指针
	//var a *int64
	//fmt.Println("var a *int64 isNil:", reflect.ValueOf(a).IsNil())
	//
	//// nil 类型
	//fmt.Println("var a *int64 isNil:", reflect.ValueOf(nil).IsValid())
	//
	//b := struct{}{}
	//// 尝试从结构体中查找"abc"字段
	//fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	//// 尝试从结构体中查找"abc"方法
	//fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	//var a int64 = 100
	// reflectSetValue(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	//reflectSetValue2(&a)
	//fmt.Println(a)

	//var a float64 = 1.0
	//reflectType(a)
	//var b int = 1
	//reflectType(b)
	//var c myInt = 1
	//reflectType(c)
	//
	//var d = person{
	//	name: "xiaoming",
	//	age:  18,
	//}
	//reflectType(d)
	//
	//var e = book{
	//	tittle: "golang",
	//}
	//reflectType(e)

	//var a float32 = 3.14
	//var b int64 = 100
	//reflectValue(a) // type is float32, value is 3.140000
	//reflectValue(b) // type is int64, value is 100
	//// 将int类型的原始值转换为reflect.Value类型
	//c := reflect.ValueOf(10)
	//fmt.Printf("type c :%T\n", c) // type c :reflect.Value

}
