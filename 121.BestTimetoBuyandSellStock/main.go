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
	fmt.Println(maxProfit1(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit1(prices))
}

// 一次遍历
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min, maxResult := math.MaxInt64, 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > maxResult {
			maxResult = price - min
		}
	}
	return maxResult
}

// 双循环
func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxResult := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxResult {
				maxResult = prices[j] - prices[i]
			}
		}
	}
	return maxResult
}
