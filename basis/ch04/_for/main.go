package main

func main() {
	/* //打印0 ～ 9
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	*/

	/*
		//计算0～100的和
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum）
	*/

	/*
		//打印乘法表
		for i := 1; i <= 9; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d*%d=%d ", i, j, i*j)
			}
			fmt.Println()
		}
	*/

	/*
		//for 循环： 使用for range, 主要用于字符串，数组，切片，map等
		str := "ABCDEFGHIJKLM哈哈哈哈"
		strrune := []rune(str)
		for i := 0; i < len(strrune); i++ {
			fmt.Printf("%c", strrune[i])
		}
		fmt.Println()
		for _, value := range str { //使用匿名变量，省略对key的输出
			//fmt.Println(key, value)
			fmt.Printf("%c\n", value)
		}

		for index := range strrune {
			fmt.Printf("%c", strrune[index])
		}
	*/

	/*
		//continue & break
		round := 0
		for {
			time.Sleep(1 * time.Second)
			round++
			if round == 5 {
				continue
			}
			fmt.Println(round)
			if round > 10 {
				break
			}
		}
	*/

}
