package main

import (
	"fmt"
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 先对每行排序
	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}
	result := 0
	// 对列循环，找出此列最大的加到 result 中
	for j := 0; j < n; j++ {
		mx := 0
		for i := 0; i < m; i++ {
			mx = max(mx, grid[i][j])
		}
		result += mx
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
	fmt.Println(deleteGreatestValue([][]int{{10}}))
}
