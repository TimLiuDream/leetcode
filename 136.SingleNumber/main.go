package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	numCount := make(map[int]int)
	for _, num := range nums {
		_, ok := numCount[num]
		if !ok {
			numCount[num] = 1
		} else {
			delete(numCount, num)
		}
	}
	for key, value := range numCount {
		if value == 1 {
			result = key
		}
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
