package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums1 := []int{-1, -100, 3, 99}
	rotate(nums, 3)
	fmt.Println(nums)
	rotate(nums1, 2)
	fmt.Println(nums1)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	nums1 = []int{-1, -100, 3, 99}
	rotate1(nums, 3)
	fmt.Println(nums)
	rotate1(nums1, 2)
	fmt.Println(nums1)
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	i := k % len(nums)
	copy(nums, append(nums[len(nums)-i:], nums[:len(nums)-i]...))
}

// 需要原地执行的话，只能用copy来复制地址了
func rotate1(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < k; i++ {
		copy(nums, append(nums[len(nums)-1:], nums[:len(nums)-1]...))
	}
}
