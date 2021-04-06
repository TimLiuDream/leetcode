package main

import "fmt"

// Label 数组
// Label 双指针

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(removeDuplicates(nums))
	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) > 2 {
		fir := 0
		sec := 2

		for sec < len(nums) {
			if nums[fir] == nums[sec] {
				nums = append(nums[:fir+1], nums[sec:]...)
			} else {
				fir++
				sec++
			}
		}
	}
	return len(nums)
}
