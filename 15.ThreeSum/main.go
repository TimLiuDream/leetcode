package main

import (
	"fmt"
	"sort"
)

// 暴力法很难解决重复三元组的问题
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for z := j + 1; z < len(nums); z++ {
				if nums[i]+nums[j]+nums[z] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[z]})
				}
			}
		}
	}
	return result
}

// 排序加双指针
func threeSum1(nums []int) [][]int {
	//先对数组排序
	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			leftValue, rightValue := nums[left], nums[right]
			if nums[i]+leftValue+rightValue > 0 {
				right--
			} else if nums[i]+leftValue+rightValue < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], leftValue, rightValue})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println("----------------")
	fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum1([]int{0, 1, 1}))
	fmt.Println(threeSum1([]int{0, 0, 0}))
}
