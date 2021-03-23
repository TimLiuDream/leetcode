package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {
	nums = append(nums, nums...)
	length := len(nums)
	ans := make([]int, length)
	var stack = make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		} else {
			ans[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return ans[0 : len(ans)/2]

}
