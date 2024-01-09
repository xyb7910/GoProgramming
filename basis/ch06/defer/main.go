package main

import "fmt"

func deferReturn() (ret int) {
	defer func() { //在return之前执行，因此有能力修改返回值
		ret++
	}()
	return 10
}

func main() {
	//连接数据库，打开文件，开始锁，无论如何，都要记得去关闭数据库，关闭文件， 解锁
	ret := deferReturn()
	fmt.Printf("ret = %d\r\n", ret)
}
