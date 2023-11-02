package main

import (
	"fmt"
	"sort"
)

func findDuplicates(nums []int) []int {
	sort.Ints(nums)
	result := make([]int, 0)
	length := len(nums)
	if length <= 1 {
		return result
	}
	left, right := 0, 1
	for left < length && right < length {
		if nums[left] == nums[right] {
			result = append(result, nums[left])
			left += 2
			right += 2
		} else {
			left++
			right++
		}
	}
	return result
}

func findDuplicates1(nums []int) []int {
	for i, _ := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	result := make([]int, 0)
	for i, num := range nums {
		if num != i+1 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates([]int{1, 1, 2}))
	fmt.Println(findDuplicates([]int{1}))

	fmt.Println(findDuplicates1([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates1([]int{1, 1, 2}))
	fmt.Println(findDuplicates1([]int{1}))
}
