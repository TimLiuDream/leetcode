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
	left, right := 0, x
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println("---------------")
	fmt.Println(mySqrt1(4))
	fmt.Println(mySqrt1(8))
	fmt.Println("--------------")
	fmt.Println(mySqrt2(4))
	fmt.Println(mySqrt2(8))
}
