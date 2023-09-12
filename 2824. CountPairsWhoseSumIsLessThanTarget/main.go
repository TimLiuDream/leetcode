package main

import (
	"fmt"
	"sort"
)

func countPairs(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}

// 排序之后加双指针
func countPairs1(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			res += right - left
			left++
		} else {
			right--
		}
	}
	return res
}

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
	fmt.Println(countPairs1([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(countPairs1([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}
