package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close(string) error
}

type writerCloser struct {
	//MyCloser // interface 也是一个类型
}

type writer struct{}

func (wc *writerCloser) Write(string) error {
	fmt.Println("write string")
	return nil
}

func (wc *writerCloser) Close(string) error {
	fmt.Println("close string")
	return nil
}

func main() {
	var mw MyWriter = &writerCloser{}
	mw.Write("hahh")
	var mc MyCloser = &writerCloser{}
	mc.Close("hahah")
}
