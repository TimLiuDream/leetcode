package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	numMap := make(map[int]struct{})
	for _, num := range nums {
		_, ok := numMap[num]
		if !ok {
			numMap[num] = struct{}{}
		} else {
			delete(numMap, num)
		}
	}
	for key, _ := range numMap {
		result = key
	}
	return result
}

func singleNumber1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

func main() {
	nums1 := []int{2, 2, 1}
	fmt.Println(singleNumber(nums1))
	fmt.Println(singleNumber1(nums1))
	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums2))
	fmt.Println(singleNumber1(nums2))
}
