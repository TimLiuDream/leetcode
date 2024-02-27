package main

import (
	"fmt"
	"sort"
)

func minimumRightShifts(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if sort.IntsAreSorted(nums) {
			return i
		}
		nums = append(nums[len(nums)-1:], nums[:len(nums)-1]...)
	}
	return -1
}

func main() {
	fmt.Println(minimumRightShifts([]int{3, 4, 5, 1, 2}))
	fmt.Println(minimumRightShifts([]int{1, 3, 5}))
	fmt.Println(minimumRightShifts([]int{2, 1, 4}))
}
