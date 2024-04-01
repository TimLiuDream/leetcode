package main

import (
	"fmt"
	"math"
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

func fib3(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}

func fib4(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

func fib6(n int) int {
	sqrt5 := math.Sqrt(5)
	p1 := math.Pow((1+sqrt5)/2, float64(n))
	p2 := math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round((p1 - p2) / sqrt5))
}
