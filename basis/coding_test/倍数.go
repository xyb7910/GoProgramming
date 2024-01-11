package main

import "fmt"

func main() {
	var a, b int64
	fmt.Scanf("%d %d", &a, &b)
	if a%b == 0 || b%a == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}

/*
input1:
6 24
output1:
Sao Multiplos

input2:
2 3
output2:
Nao sao Multiplos
*/
