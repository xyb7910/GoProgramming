package ch11

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect: %d, actual: %d", 4, re)
	}
}

func TestAdd2(t *testing.T) {
	fmt.Println("yes1")
	if testing.Short() {
		t.Skip("short 模式下跳过")
	}
	fmt.Println("yes")

	res := add(1, 5)
	if res != 6 {
		t.Errorf("expect: %d, actual: %d", 6, res)
	}
}

func TestAdd3(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, 0, -1},
		{78, 2, 80},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect: %d, actual: %d", value.out, re)
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 122
	b = 111
	c = 233
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d + %d, expect: %d, actual: %d", a, b, c, actual)
		}
	}
}
