package main

import "fmt"

func singleNumber(nums []int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	result := make([]int, 0)
	for num, count := range numMap {
		if count == 1 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{-1, 0}))
	fmt.Println(singleNumber([]int{0, 1}))
}
