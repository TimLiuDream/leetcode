package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{0}))
	fmt.Println("=======================")
	fmt.Println(missingNumber1([]int{3, 0, 1}))
	fmt.Println(missingNumber1([]int{0, 1}))
	fmt.Println(missingNumber1([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber1([]int{0}))
	fmt.Println("=======================")
	fmt.Println(missingNumber2([]int{3, 0, 1}))
	fmt.Println(missingNumber2([]int{0, 1}))
	fmt.Println(missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber2([]int{0}))
	fmt.Println("=======================")
}

func missingNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

func missingNumber1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)+1; i++ {
		if i != nums[i] {
			return i
		}
		if i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}

func missingNumber2(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}
