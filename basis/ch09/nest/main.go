package main

import "fmt"

// 接口定义的嵌套
type MyReader interface {
	Read() string
}

type MyWriter interface {
	Writer() string
}

type MyReadWriter interface {
	MyReader
	MyWriter
	ReadWrit()
}

type SreadWriter struct {
}

func (s *SreadWriter) Writer() string {
	//TODO implement me
	fmt.Println("read writer")
	return ""
}

func main() {
	var mw MyWriter = &SreadWriter{}
	mw.Writer()
}
