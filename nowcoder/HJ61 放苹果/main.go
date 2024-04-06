package main

import "fmt"

func countWays(m, n int) int {
	// 边界情况
	if m == 0 || n == 1 {
		return 1
	}
	if m < 0 || n < 1 {
		return 0
	}

	// 分两种情况：1. 至少有一个盘子空着；2. 所有盘子都有苹果
	return countWays(m, n-1) + countWays(m-n, n)
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	ways := countWays(m, n)
	fmt.Println(ways)
}
