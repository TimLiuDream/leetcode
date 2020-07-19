package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums1))
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2))
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums3))
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	m := make(map[int]bool)
	for _, v := range nums {
		if _, found := m[v]; !found {
			m[v] = true
		} else {
			return true
		}
	}
	return false
}
