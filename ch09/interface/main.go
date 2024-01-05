package main

import "fmt"

//接口的定义
type Duck interface {
	//方法申请
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("gaga~")
}
func (pd *pskDuck) Walk() {
	fmt.Println("this is pskduck walk~")
}
func (pd *pskDuck) Swimming() {

	fmt.Println("this is pskduck swimming~")
}

func main() {
	/*
		到处是interface， 到处是鸭子
		当看到一个鸟走起来很像鸭子，游泳起来很像鸭子，叫起来很像鸭子，那么这只鸟就是鸭子
		动词 方法 具备某些方法 强调事物的外部行为，而不是内部结构
	*/

	var d Duck = &pskDuck{}
	d.Walk()
}
