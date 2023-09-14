package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib2(4))
	fmt.Println(fib2(5))
}

// 递归不缓存
func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// 递归缓存
func fib1(N int) int {
	m := make(map[int]int)
	var f func(n int) int
	f = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		if v, ok := m[n]; ok {
			return v
		}
		m[n] = f(n-1) + f(n-2)
		return m[n]
	}
	return f(N-1) + f(N-2)
}

// 迭代不缓存
func fib2(N int) int {
	if N <= 1 {
		return N
	}

	nums := make([]int, N+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= N; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[N]
}
