package main

import (
	"fmt"
)

// Label 位运算

// solution https://leetcode-cn.com/problems/power-of-four/solution/4de-mi-by-leetcode-solution-b3ya/

func main() {
	n := 2
	fmt.Println(isPowerOfFour(n))
	n = 5
	fmt.Println(isPowerOfFour(n))
	n = 1
	fmt.Println(isPowerOfFour(n))
	n = 64
	fmt.Println(isPowerOfFour(n))
}

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
