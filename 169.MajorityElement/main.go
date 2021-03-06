package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]//因为题目要求众数是说元素的个数大于n/2的
}

func majorityElement1(nums []int) int {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		count, ok := numCountMap[num]
		if !ok {
			numCountMap[num] = 1
		} else {
			numCountMap[num] = count + 1
		}
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

func main() {
	nums := []int{3, 2, 3}
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement(nums1))

	fmt.Println(majorityElement1(nums))
	fmt.Println(majorityElement1(nums1))
}
