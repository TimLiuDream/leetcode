package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5}))
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3}))
}

// 在第 i 次翻转之后，我们希望 [1,i] 内的所有位都是 1 ，这等价于「前 i 次翻转中下标的最大值等于 i」。
func numTimesAllBlue(flips []int) int {
	n, ans, right := len(flips), 0, 0
	for i := 0; i < n; i++ {
		if flips[i] > right {
			right = flips[i]
		}
		if right == i+1 {
			ans++
		}
	}
	return ans
}
