package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_CreateFile(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", file)
	}
	defer file.Close()
}

func Test_CreateDir(t *testing.T) {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_CreateDirAll(t *testing.T) {
	err := os.MkdirAll("ms/one/two", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func Test_RemoveFile(t *testing.T) {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func Test_RemoveDir(t *testing.T) {
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func Test_RemoveDirAll(t *testing.T) {
	err := os.RemoveAll("ms/one/two")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
