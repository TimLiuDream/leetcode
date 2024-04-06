package main

import (
	"fmt"
)

func main() {
	var n int32
	fmt.Scan(&n)
	num := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			num++
		}
	}
	fmt.Print(num)
}
