package main

import "fmt"

func main() {
	fmt.Println(findChampion([][]int{{0, 1}, {0, 0}}))
	fmt.Println(findChampion([][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}}))
}

func findChampion(grid [][]int) int {
	mw := make(map[int]struct{})
	ml := make(map[int]struct{})
	for i, rows := range grid {
		for j, col := range rows {
			if col == 1 {
				mw[i] = struct{}{}
				ml[j] = struct{}{}
			}
		}
	}
	ans := 0
	for i, _ := range mw {
		if _, ok := ml[i]; !ok {
			return i
		}
	}
	return ans
}

// 方法一：每行求和
func findChampion1(grid [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		sum := 0
		for _, val := range grid[i] {
			sum += val
		}
		if sum == n-1 {
			return i
		}
	}
	return -1
}
