package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	fmt.Println(maxSumDivThree([]int{4}))
	fmt.Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxSumDivThree(nums []int) int {
	sort.Ints(nums)
	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}
	allResult, remove := sum(nums), math.MaxInt8
	if allResult%3 == 0 {
		remove = 0
	} else if allResult%3 == 1 {
		if len(v[1]) > 0 {
			remove = min(remove, v[1][0])
		}
		if len(v[2]) > 1 {
			remove = min(remove, v[2][0]+v[2][1])
		}
	} else {
		if len(v[1]) > 1 {
			remove = min(remove, v[1][0]+v[1][1])
		}
		if len(v[2]) > 0 {
			remove = min(remove, v[2][0])
		}
	}
	return allResult - remove
}
