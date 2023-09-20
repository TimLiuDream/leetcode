package main

import "fmt"

func minCount(coins []int) int {
	count := 0
	for _, coin := range coins {
		if coin%2 == 1 {
			count += coin/2 + 1
		} else {
			count += coin / 2
		}
	}
	return count
}

func minCount1(coins []int) int {
	sum := 0
	for _, coin := range coins {
		// 假设这次是 2 的整数倍，那么 (coin+1)/2 == coin/2
		// 假设是 2n+1，那么 (coin+1)/2 == coin/2 +1
		sum += (coin + 1) / 2
	}
	return sum
}

func main() {
	fmt.Println(minCount([]int{4, 2, 1}))
	fmt.Println(minCount([]int{2, 3, 10}))
}
