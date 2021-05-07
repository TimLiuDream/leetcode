package main

import "fmt"

// Label 位运算
// Label 数组

func main() {
	n := 5
	start := 0
	fmt.Println(xorOperation(n, start))

	n = 4
	start = 3
	fmt.Println(xorOperation(n, start))

	n = 1
	start = 7
	fmt.Println(xorOperation(n, start))

	n = 10
	start = 5
	fmt.Println(xorOperation(n, start))
}

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		result = result ^ (start + i*2)
	}
	return result
}
