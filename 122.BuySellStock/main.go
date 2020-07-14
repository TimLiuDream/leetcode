package main

import "fmt"

// 所有的上升线段都计算进去
func maxProfit(prices []int) int {
	result := 0
	for index := 0; index < len(prices)-1; index++ {
		if prices[index+1] > prices[index] {
			result += prices[index+1] - prices[index]
		}
	}
	return result
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(nums))
	fmt.Println(maxProfit(nums1))
	fmt.Println(maxProfit(nums2))
}
