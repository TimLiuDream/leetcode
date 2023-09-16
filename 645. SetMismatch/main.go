package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	res := []int{0, 0}
	for i := 1; i <= len(nums); i++ {
		count, ok := m[i]
		if ok && count == 2 {
			res[0] = i
		}
		if !ok {
			res[1] = i
		}
	}
	return res
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{3, 2, 3, 4, 6, 5}))
}
