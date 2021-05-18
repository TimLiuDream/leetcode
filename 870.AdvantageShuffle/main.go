package main

import (
	"fmt"
	"sort"
)

// Label 贪心算法
// Label 数组

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{1, 10, 4, 11}
	fmt.Println(advantageCount(nums1, nums2))

	nums1 = []int{12, 24, 8, 32}
	nums2 = []int{13, 25, 32, 11}
	fmt.Println(advantageCount(nums1, nums2))
}

func advantageCount(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	sort.Ints(nums1)
	for i, num2 := range nums2 {
		for j, num1 := range nums1 {
			if num1 > num2 {
				result = append(result, num1)
				if j < len(nums1) {
					nums1 = append(nums1[:j], nums1[j+1:]...)
				}
				break
			}
		}
		if len(result) == i {
			result = append(result, nums1[0])
			nums1 = nums1[1:]
		}
	}
	return result
}
