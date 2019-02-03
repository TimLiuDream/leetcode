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

func main() {
	nums1 := []int{2, 2, 1}
	fmt.Println(singleNumber(nums1))
	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums2))
}
