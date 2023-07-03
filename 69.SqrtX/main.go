package main

import (
	"fmt"
	"math"
)

// 使用api来求解
func mySqrt(x int) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(ff)
}

// 使用牛顿法求平方根
func mySqrt1(x int) int {
	res := x
	//牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

// 二分查找
func mySqrt2(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))

	fmt.Println(mySqrt1(4))
	fmt.Println(mySqrt1(8))

	fmt.Println(mySqrt2(4))
	fmt.Println(mySqrt2(8))
}
