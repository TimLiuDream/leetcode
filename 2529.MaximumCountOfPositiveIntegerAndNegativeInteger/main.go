package main

import "fmt"

func maximumCount(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		if num < 0 {
			a++
			continue
		}
		if num > 0 {
			b++
			continue
		}
	}
	ans := max(a, b)
	return ans
}

func maximumCount1(nums []int) int {
	pos1 := lowerBound(nums, 0)
	pos2 := lowerBound(nums, 1)
	return max(pos1, len(nums)-pos2)
}

func lowerBound(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] >= val {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
	fmt.Println(maximumCount([]int{-3, -2, -1, 0, 0, 1, 2}))
	fmt.Println(maximumCount([]int{5, 20, 66, 1314}))

	fmt.Println(maximumCount1([]int{-2, -1, -1, 1, 2, 3}))
	fmt.Println(maximumCount1([]int{-3, -2, -1, 0, 0, 1, 2}))
	fmt.Println(maximumCount1([]int{5, 20, 66, 1314}))
}
