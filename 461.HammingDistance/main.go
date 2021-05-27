package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

// Label 位运算

// solution https://leetcode-cn.com/problems/hamming-distance/solution/yi-ming-ju-chi-by-leetcode-solution-u1w7/

func main() {
	x := 0
	y := 1
	fmt.Println(hammingDistance(x, y))
}

func hammingDistance(x int, y int) int {
	xStr := convertToBin(x)
	yStr := convertToBin(y)

	count := 0
	for i := 0; i < 32; i++ {
		if xStr[i] != yStr[i] {
			count++
		}
	}
	return count
}

func convertToBin(num int) string {
	c := "00000000000000000000000000000000"
	s := ""

	if num == 0 {
		return c
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}

	if len(s) <= 32 {
		s = string(c[:32-len(s)]) + s
	}
	return s
}

func hammingDistance1(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
