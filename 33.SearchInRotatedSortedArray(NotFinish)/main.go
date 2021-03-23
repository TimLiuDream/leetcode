package main

import "fmt"

func search(nums []int, target int) int {
	result := 0
	nums1 := nums[:len(nums)/2]
	nums2 := nums[len(nums)/2:]
	if nums1[0] > nums1[len(nums1)/2] {
		search(nums1, target)
	} else if nums1[0] <= nums1[len(nums1)/2] {
		// 二分查找nums1
		ok, index := binarySearch(nums1, target)
		if ok {
			result = index
		}
	}
	if nums2[0] > nums2[len(nums2)/2] {
		search(nums2, target)
	} else {
		// 二分查找nums2
		ok, index := binarySearch(nums2, target)
		if ok {
			result = index
		}
	}
	return result
}

func binarySearch(nums []int, target int) (found bool, result int) {
	if nums[0] > target {
		return
	}
	minIndex := 0
	maxIndex := len(nums)
	for minIndex <= maxIndex {
		result = (minIndex + maxIndex) / 2
		if nums[result] == target {
			found = true
			break
		} else if nums[result] > target {
			minIndex++
		} else {
			maxIndex++
		}
	}
	return
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}
