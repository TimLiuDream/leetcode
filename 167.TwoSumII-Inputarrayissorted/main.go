package main

import "fmt"

// Label 数组
// Label 双指针
// Label 二分查找

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))

	numbers = []int{2, 3, 4}
	target = 6
	fmt.Println(twoSum(numbers, target))

	numbers = []int{-1, 0}
	target = -1
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	var result []int
	pre := 0
	next := len(numbers) - 1
	for pre < next {
		if numbers[pre]+numbers[next] == target {
			result = []int{pre + 1, next + 1}
			break
		} else if numbers[pre]+numbers[next] > target {
			next--
		} else {
			pre++
		}
	}
	return result
}
