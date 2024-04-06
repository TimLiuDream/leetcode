package main

import (
	"fmt"
)

func main() {
	//var money, count int
	//fmt.Scan(&money, &count)
	//goods := make([][]int, count)
	//for i := 0; i < count; i++ {
	//	goods[i] = make([]int, 3)
	//	fmt.Scan(&goods[i][0], &goods[i][1], &goods[i][2])
	//}
	cal(1000, 5, [][]int{{800, 2, 0}, {400, 5, 1}, {300, 5, 1}, {400, 3, 0}, {500, 2, 0}})
}

func cal(money, count int, goods [][]int) {
	costs := make([][]int, 0)
	prices := make([][]int, 0)
	for i := 0; i < count; i++ {
		if goods[i][2] != 0 {
			continue
		}
		cost := make([]int, 0)
		price := make([]int, 0)
		cost = append(cost, goods[i][0]*goods[i][1])
		price = append(price, goods[i][0])
		for j := 0; j < count; j++ {
			if goods[j][2]-1 == i {
				cost = append(cost, goods[j][0]*goods[j][1]+cost[0])
				price = append(price, goods[j][0]+price[0])
			}
		}
		if len(cost) == 3 {
			cost = append(cost, cost[1]+cost[2]-cost[0])
			price = append(price, price[1]+price[2]-price[0])
		}
		costs = append(costs, cost)
		prices = append(prices, price)
	}
	size := len(costs)
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, money+1)
		dp[i][0] = 0
	}
	for j := 0; j < money+1; j += 10 {
		dp[0][j] = 0
	}
	for i := 1; i < size+1; i++ {
		for j := 10; j < money+1; j += 10 {
			max := dp[i-1][j]
			for k, v := range prices[i-1] {
				if j-v >= 0 {
					if max < dp[i-1][j-v]+costs[i-1][k] {
						max = dp[i-1][j-v] + costs[i-1][k]
					}
				}
			}
			dp[i][j] = max
		}
	}
	fmt.Println(dp[size][money])
}
