package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// nums := []int{1}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	i := k % len(nums)
	copy(nums, append(nums[len(nums)-i:], nums[:len(nums)-i]...))
}
