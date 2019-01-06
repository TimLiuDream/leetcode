package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {//这是一个返回条件
		return 1
	}
	if n < 0 {//当n为负数时，那么就返回以-n为幂的倒数
		return 1 / myPow(x, -n)
	}
	y := myPow(x, n/2)//使用分治，不断的以n/2来计算
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(34.00515, -3))
}
