package main

import (
	"fmt"
	"math"
)

//使用api来求解
func mySqrt(x int) int {
	f := float64(x)
	ff := math.Sqrt(f)
	return int(ff)
}

//使用牛顿法求平方根
func mySqrt1(x int) int {
	res := x
	//牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func main() {
	fmt.Println(mySqrt1(4))
	fmt.Println(mySqrt1(8))
}
