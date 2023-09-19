package main

import (
	"fmt"
)

// dfs
func numIslands(grid [][]byte) int {
	result := 0
	for i, bytes := range grid {
		for j, b := range bytes {
			if b == '1' {
				result += 1
				dfs(grid, i, j)
			}
		}
	}
	return result
}

func dfs(grid [][]byte, r, c int) {
	rc, cc := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= rc || c >= cc || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
