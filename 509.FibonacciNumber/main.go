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
func fib1(n int) int {
	m := make(map[int]int)
	var f func(n int) int
	f = func(i int) int {
		if i <= 1 {
			return i
		}
		v, ok := m[i]
		if ok {
			return v
		}
		value := f(i-1) + f(i-2)
		m[i] = value
		return value
	}
	return f(n-1) + f(n-2)
}

// 迭代不缓存
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		b, a = a+b, b
	}

	return b
}
