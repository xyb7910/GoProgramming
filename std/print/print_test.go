package print

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Print("我是控制台打印，我不换行 可以自己控制换行")
}
func TestPrintln(t *testing.T) {
	fmt.Println("我是控制台打印，我换行")
}
func TestPrintf(t *testing.T) {
	fmt.Printf("我是控制台打印，%s \n", "这是格式化占位符信息,可以自己控制换行")
}

type User struct {
	Id int64
}

func TestPrintf1(t *testing.T) {
	user := &User{Id: 1}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%%\n")
}

func TestPrintf2(t *testing.T) {
	fmt.Printf("%t\n", true)
}

func TestPrintf3(t *testing.T) {
	n := 180
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n", n)
	a := 96
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", 0x4E2D)
}

func TestPrintf4(t *testing.T) {
	f := 18.54
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}

func TestPrintf5(t *testing.T) {
	s := "我是字符串"
	str := "高乐乐"
	b := []byte{65, 66, 67}
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", str)
	fmt.Printf("%X\n", s)
}

func TestPrintf6(t *testing.T) {
	n := 13.14
	fmt.Printf("%f\n", n)
	fmt.Printf("%10f\n", n)
	fmt.Printf("%10s\n", "我是字符串")
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%10.2f\n", n)
	fmt.Printf("%10.f\n", n)
}

func TestPrintf7(t *testing.T) {
	s := "我是字符串"
	fmt.Printf("%  d\n", 10)
	fmt.Printf("%s\n", s)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%10.2f\n", 10.14)
	fmt.Printf("%-10.2f\n", 10.14)
	fmt.Printf("%010s\n", s)
}

func TestFPrint1(t *testing.T) {
	file, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	fmt.Fprintln(file, "追加写入")
	file.Close()
}

func TestSPrint(t *testing.T) {
	s1 := fmt.Sprint("张三")
	name := "张三"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("张三")
	fmt.Println(s1, s2, s3)
}

func TestErrorf(t *testing.T) {
	err := fmt.Errorf("用户名格式不正确：%s", "@#￥哈哈")
	if err != nil {
		panic(err)
	}
}
