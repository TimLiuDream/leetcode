package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2] //因为题目要求众数是说元素的个数大于n/2的
}

func majorityElement1(nums []int) int {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}

	maxCount := 0
	maxKey := 0
	for key, value := range numCountMap {
		if value >= maxCount {
			maxKey = key
		}
	}
	return maxKey
}

func majorityElement2(nums []int) int {
	result, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			result = num
			count++
		} else if result != num {
			count--
		} else {
			count++
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 3}
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement(nums1))

	fmt.Println(majorityElement1(nums))
	fmt.Println(majorityElement1(nums1))

	fmt.Println(majorityElement2(nums))
	fmt.Println(majorityElement2(nums1))
}
