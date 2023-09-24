package main

import "fmt"

// 使用模拟方式来迭代
func islandPerimeter(grid [][]int) int {
	type pair struct{ x, y int }
	pairs := []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := 0

	for x, rValue := range grid {
		for y, cValue := range rValue {
			if cValue != 1 {
				continue
			}
			for _, p := range pairs {
				rx, ry := x+p.x, y+p.y
				if (rx < 0 || rx >= len(grid)) || (ry < 0 || ry >= len(rValue)) || grid[rx][ry] == 0 {
					result++
				}
			}
		}
	}

	return result
}

// dfs
func islandPerimeter1(grid [][]int) int {
	type pair struct{ x, y int }
	pairs := []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := 0
	rLen, cLen := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rLen || y < 0 || y >= cLen || grid[x][y] == 0 {
			result++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, p := range pairs {
			dfs(x+p.x, y+p.y)
		}
	}
	for x, row := range grid {
		for y, col := range row {
			if col == 1 {
				dfs(x, y)
			}
		}
	}
	return result
}

// bfs
func islandPerimeter2(grid [][]int) int {
	type pair struct{ x, y int }
	pairs := []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := 0
	rLen, cLen := len(grid), len(grid[0])

	queue := []pair{}
	visited := make([][]bool, rLen)
	for i, _ := range visited {
		visited[i] = make([]bool, cLen)
	}

	for x, row := range grid {
		for y, col := range row {
			if col == 1 {
				queue = append(queue, pair{x, y})
				visited[x][y] = true
			}
		}
	}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, p := range pairs {
			newX, newY := curr.x+p.x, curr.y+p.y
			if newX < 0 || newX >= rLen || newY < 0 || newY >= cLen || grid[newX][newY] == 0 {
				result++
			} else if !visited[newX][newY] {
				queue = append(queue, pair{x: newX, y: newY})
				visited[newX][newY] = true
			}
		}
	}

	return result
}

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter([][]int{{1}}))
	fmt.Println(islandPerimeter([][]int{{1, 0}}))
	fmt.Println("-----------------------------------")
	fmt.Println(islandPerimeter1([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter1([][]int{{1}}))
	fmt.Println(islandPerimeter1([][]int{{1, 0}}))
	fmt.Println("-----------------------------------")
	fmt.Println(islandPerimeter2([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter2([][]int{{1}}))
	fmt.Println(islandPerimeter2([][]int{{1, 0}}))
}
