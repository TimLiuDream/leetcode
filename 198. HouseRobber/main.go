package main

import (
	"fmt"
)

func rob(nums []int) int {
	// 因为当一间房间都没有的时候，肯定没有钱可以偷，
	// 当房间只有一件的时候，只能偷那一间
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 0, nums[0]
	if len(nums) <= 1 {
		return dp[len(nums)]
	}
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(dp)-1]
}

func max(i, j int) int {
	m := i
	if j > i {
		m = j
	}
	return m
}

func main() {
	fmt.Println(rob([]int{0}))
	fmt.Println(rob([]int{0, 1}))
	fmt.Println(rob([]int{0, 1, 2}))
	fmt.Println(rob([]int{1, 3, 3, 2}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 7, 10}))
}
