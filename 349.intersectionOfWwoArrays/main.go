package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection1(nums1, nums2))
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersection1(nums1, nums2))
	nums1 = []int{54, 93, 21, 73, 84, 60, 18, 62, 59, 89, 83, 89, 25, 39, 41, 55, 78, 27, 65, 82, 94, 61, 12, 38, 76, 5, 35, 6, 51, 48, 61, 0, 47, 60, 84, 9, 13, 28, 38, 21, 55, 37, 4, 67, 64, 86, 45, 33, 41}
	nums2 = []int{17, 17, 87, 98, 18, 53, 2, 69, 74, 73, 20, 85, 59, 89, 84, 91, 84, 34, 44, 48, 20, 42, 68, 84, 8, 54, 66, 62, 69, 52, 67, 27, 87, 49, 92, 14, 92, 53, 22, 90, 60, 14, 8, 71, 0, 61, 94, 1, 22, 84, 10, 55, 55, 60, 98, 76, 27, 35, 84, 28, 4, 2, 9, 44, 86, 12, 17, 89, 35, 68, 17, 41, 21, 65, 59, 86, 42, 53, 0, 33, 80, 20}
	fmt.Println(intersection1(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, num := range nums1 {
		m[num] = struct{}{}
	}
	result := make([]int, 0)
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			result = append(result, num)
			delete(m, num)
		}
	}
	return result
}

func intersection1(nums1, nums2 []int) []int {
	// 排序
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 双指针
	result, pre := make([]int, 0), math.MinInt8
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		iV, jV := nums1[i], nums2[j]
		if iV != jV {
			if iV < jV {
				i++
			} else {
				j++
			}
		} else {
			if iV != pre {
				result = append(result, iV)
				pre = iV
			}
			i++
			j++
		}
	}
	return result
}
