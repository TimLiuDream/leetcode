package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums1))
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2))
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums3))

	fmt.Println(containsDuplicate1(nums1))
	fmt.Println(containsDuplicate1(nums2))
	fmt.Println(containsDuplicate1(nums3))
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	m := make(map[int]struct{})
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
