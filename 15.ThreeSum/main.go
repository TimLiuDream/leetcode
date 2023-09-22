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
		j := i + 1         // 左指针
		z := len(nums) - 1 // 右指针
		for z > j {
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}
				j++
				z--
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
