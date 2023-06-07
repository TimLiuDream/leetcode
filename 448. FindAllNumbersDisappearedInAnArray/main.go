package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
	fmt.Println("=======")
	fmt.Println(findDisappearedNumbers1([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers1([]int{1, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	raw := make([]int, len(nums))
	for _, num := range nums {
		if raw[num-1] == 0 {
			raw[num-1] = num
		}
	}
	result := make([]int, 0)
	for index, n := range raw {
		if n == 0 {
			result = append(result, index+1)
		}
	}
	return result
}

func findDisappearedNumbers1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		num := abs(nums[i])
		if nums[num-1] > 0 {
			nums[num-1] = -1 * nums[num-1]
		}
	}
	result := make([]int, 0)
	for index, num := range nums {
		if num > 0 {
			result = append(result, index+1)
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}
