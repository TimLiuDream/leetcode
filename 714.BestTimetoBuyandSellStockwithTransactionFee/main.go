package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	// prices := []int{1, 3, 7, 5, 10, 3}
	// fee := 3

	fmt.Println(maxProfit(prices, fee))
}

// func maxProfit(prices []int, fee int) int {
// 	result := 0
// 	i := 0
// 	for index := 1; index < len(prices)-1; index++ {
// 		if prices[index]-prices[i] > fee {
// 			result += prices[index] - prices[i] - fee
// 			i = index + 1
// 			index++
// 		}
// 	}
// 	return result
// }

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
