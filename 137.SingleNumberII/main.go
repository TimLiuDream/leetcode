package main

import "fmt"

// Label 位运算

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
	nums = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
	nums = []int{30000,500,100,30000,100,30000,100}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for num, occ := range freq {
		if occ == 1 {
			return num
		}
	}
	return 0
}
