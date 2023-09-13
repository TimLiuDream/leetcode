package main

import "fmt"

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid)
	// 构造骑士的路径
	path := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[grid[i][j]][0] = i
			path[grid[i][j]][1] = j
		}
	}
	// 判断当前位置和上一次的位置是否符合：
	// x2 - x1 = 1 || x2 - x1 = 2
	// y2 - y1 = 1 || y2 - y1 = 2
	for i := 1; i < n*n; i++ {
		x := abs(path[i][0] - path[i-1][0])
		y := abs(path[i][1] - path[i-1][1])
		if x*y != 2 {
			return false
		}
	}
	return true
}

func checkValidGridDFS(grid [][]int) bool {
	n := len(grid)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(grid, 0, 0, visited)
}

func dfs(grid [][]int, x, y int, visited [][]bool) bool {
	n := len(grid)

	// 越界或已访问过的格子
	if x < 0 || y < 0 || x >= n || y >= n || visited[x][y] {
		return false
	}

	// 到达目标格子
	if x == n-1 && y == n-1 {
		return true
	}

	// 标记当前格子为已访问
	visited[x][y] = true

	// 往四个方向进行深度优先搜索
	dx := []int{1, 1, -1, -1, 2, 2, -2, -2}
	dy := []int{2, -2, 2, -2, 1, -1, 1, -1}
	for i := 0; i < 8; i++ {
		nextX := x + dx[i]
		nextY := y + dy[i]
		if dfs(grid, nextX, nextY, visited) {
			return true
		}
	}

	// 没有找到有效路径，回溯
	visited[x][y] = false
	return false
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func checkValidGridBFS(grid [][]int) bool {
	n := len(grid)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	queue := [][]int{{0, 0}}
	visited[0][0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		x := cur[0]
		y := cur[1]

		// 到达目标格子
		if x == n-1 && y == n-1 {
			return true
		}

		// 往四个方向进行广度优先搜索
		dx := []int{1, 1, -1, -1, 2, 2, -2, -2}
		dy := []int{2, -2, 2, -2, 1, -1, 1, -1}
		for i := 0; i < 8; i++ {
			nextX := x + dx[i]
			nextY := y + dy[i]
			if nextX >= 0 && nextX < n && nextY >= 0 && nextY < n && !visited[nextX][nextY] {
				queue = append(queue, []int{nextX, nextY})
				visited[nextX][nextY] = true
			}
		}
	}

	return false
}

func checkValidGridDP(grid [][]int) bool {
	n := len(grid)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	dp[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && dp[i-1][j] && abs(grid[i][j]-grid[i-1][j]) == 2 {
				dp[i][j] = true
			}
			if j > 0 && dp[i][j-1] && abs(grid[i][j]-grid[i][j-1]) == 2 {
				dp[i][j] = true
			}
		}
	}

	return dp[n-1][n-1]
}

func main() {
	grid := [][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}
	fmt.Println(checkValidGrid(grid))
	grid1 := [][]int{{0, 3, 6}, {5, 8, 1}, {2, 7, 4}}
	fmt.Println(checkValidGrid(grid1))
}
