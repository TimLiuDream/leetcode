package main

import (
	"fmt"
	"math"
)

// Label 数组
// Label 动态规划

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := math.MaxInt64
	maxResult := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > maxResult {
			maxResult = prices[i] - min
		}
	}
	return maxResult
}
