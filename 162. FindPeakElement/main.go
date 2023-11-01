package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	length := len(nums)
	nums = append([]int{math.MinInt64}, append(nums, math.MinInt64)...)
	index := 1
	for ; index <= length; index++ {
		if nums[index] > nums[index-1] && nums[index] > nums[index+1] {
			break
		}
	}
	return index - 1
}

func findPeakElement1(nums []int) int {
	result, max := 0, math.MinInt64
	for i, num := range nums {
		if num > max {
			result = i
			max = num
		}
	}
	return result
}

func findPeakElement2(nums []int) int {
	n := len(nums)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))

	fmt.Println(findPeakElement1([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement1([]int{1, 2, 1, 3, 5, 6, 4}))

	fmt.Println(findPeakElement2([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement2([]int{1, 2, 1, 3, 5, 6, 4}))
}
