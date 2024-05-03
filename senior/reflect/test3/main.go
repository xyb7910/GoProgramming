package main

import (
	"fmt"
	"reflect"
)

// User 用户信息
type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	emails  []string `json:"emails"`
	address string   `json:"address"`
}

// GetUser 获取用户信息
func (u User) GetUser() (string, int, []string, string) {
	return u.Name, u.Age, u.emails, u.address
}

// SetUser 设置用户信息
func (u User) SetUser(Name string, Age int, emails []string, address string) {
	u.Name = Name
	u.Age = Age
	u.emails = emails
	u.address = address
}

// PrintUser 打印用户信息
func (u User) PrintUser() {
	fmt.Println(u.Name, u.Age, u.emails, u.address)
}

func main() {
	u := User{
		Name:    "xiaoming",
		Age:     18,
		emails:  []string{"test@eamil.com", "test@qq.com"},
		address: "beijing",
	}

	//PrintStructField(u)
	//PrintStructFn(u)
	ReflectChangeStruct(&u)
}

// PrintStructField 打印结构体字段
func PrintStructField(v interface{}) {
	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	//1、通过类型变量里面的 Field 可以获取结构体的字段
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field %d:  Name:%s Type:%s JsonTag:%s\n",
			i, t.Field(i).Name, t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}

	//2、通过类型变量里面的 FieldByName 可以获取结构体的字段
	field, ok := t.FieldByName("Name")
	if ok {
		fmt.Println(field.Name)
	} else {
		fmt.Println("not found")
	}
}

// PrintStructFn 打印结构体方法
func PrintStructFn(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("not struct")
		return
	}
	//1、通过类型变量里面的 Method 可以获取结构体的方法
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}
	//2、通过类型变量获取这个结构体有多少个方法
	fmt.Println(t.NumMethod())
	//3、执行方法 (注意需要使用值变量，并且要注意参数)
	val.MethodByName("PrintUser").Call(nil)
	//4、执行方法传入参数 (注意需要使用值变量，并且要注意参数)
	var params []reflect.Value // 参数
	params = append(params, reflect.ValueOf("xiaoming"))
	params = append(params, reflect.ValueOf(18))
	params = append(params, reflect.ValueOf([]string{"test@eamil.com", "test@qq.com"}))
	params = append(params, reflect.ValueOf("beijing"))
	fmt.Println(val.MethodByName("SetUser").Call(params))

	// 5、执行方法获取方法的值
	info := val.MethodByName("GetUser").Call(nil)
	fmt.Println(info)
}

// ReflectChangeStruct 通过反射修改结构体字段
func ReflectChangeStruct(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	name := val.Elem().FieldByName("Name")
	name.SetString("xiaoyan") // 设置字段值

	age := val.Elem().FieldByName("Age")
	age.SetInt(20) // 设置字段值
}
