package main

import (
	"fmt"
	"sort"
)

// Label 数组
// Label 二分查找

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	fmt.Println(shipWithinDays(weights, D))

	weights = []int{3, 2, 2, 4, 1, 4}
	D = 3
	fmt.Println(shipWithinDays(weights, D))
}

func shipWithinDays(weights []int, D int) int {
	// 确定二分查找左右边界
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1 // 需要运送的天数
		sum := 0 // 当前这一天已经运送的包裹重量之和
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}
