package main

import "fmt"

// Label 动态规划

func main() {
	steps := 3
	arrLen := 2
	fmt.Println(numWays(steps, arrLen))
}

func numWays(steps, arrLen int) int {
	const mod = 1e9 + 7
	maxColumn := min(arrLen-1, steps)
	dp := make([]int, maxColumn+1)
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		dpNext := make([]int, maxColumn+1)
		for j := 0; j <= maxColumn; j++ {
			dpNext[j] = dp[j]
			if j-1 >= 0 {
				dpNext[j] = (dpNext[j] + dp[j-1]) % mod
			}
			if j+1 <= maxColumn {
				dpNext[j] = (dpNext[j] + dp[j+1]) % mod
			}
		}
		dp = dpNext
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
