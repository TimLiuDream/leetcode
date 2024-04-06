package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a, b := 1, 1
	if n <= 2 {
		fmt.Println(1)
		return
	}
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
}
