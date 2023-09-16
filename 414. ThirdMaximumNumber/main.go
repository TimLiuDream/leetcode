package main

import (
	"fmt"
	"math"
	"sort"
)

// 使用计数器
func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, counter := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			counter++
			if counter == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}

// 三个变量
func thirdMax1(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if a > num && num > b {
			b, c = num, b
		} else if b > num && num > c {
			c = num
		}
	}
	if c != math.MinInt64 {
		return c
	}
	return a
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{1, 1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))

	fmt.Println(thirdMax1([]int{3, 2, 1}))
	fmt.Println(thirdMax1([]int{1, 2}))
	fmt.Println(thirdMax1([]int{1, 1, 2}))
	fmt.Println(thirdMax1([]int{2, 2, 3, 1}))
}
