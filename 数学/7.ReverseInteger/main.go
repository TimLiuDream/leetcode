package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	//用来记录原数据是正还是负的
	sign := 1

	//处理负数
	if x < 0 {
		sign = -1
		x = sign * x
	}

	result := 0
	for x > 0 {
		//弹出最后一位
		temp := x % 10
		//把temp放到result的最前面一位
		result = result*10 + temp
		//去掉x的最后一位
		x = x / 10
	}

	//还原正负数
	result = sign * result

	//解决溢出问题
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(-123123457687980))
}
