package main

import "fmt"

var a, b, c, d int64

func main() {
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)
	fmt.Printf("DIFERENCA = %d\n", a*b-c*d)
}

/*
input :
5
6
7
8
output:
DIFERENCA = -26
*/
