package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}

// 不缓存
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

// 缓存
func fib2(N int) int {
	if N <= 1 {
		return N
	}

	nums := make([]int, 0)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= N; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[N]
}
