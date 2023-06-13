package main

import "fmt"

func main() {
	nums := []int{4, 4, 2, 4, 3}
	fmt.Println(unequalTriplets(nums))

	nums1 := []int{1, 1, 1, 1, 1}
	fmt.Println(unequalTriplets(nums1))
}

func unequalTriplets(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}
