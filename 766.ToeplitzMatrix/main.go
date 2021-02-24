package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
	matrix = [][]int{{1, 2}, {2, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for row := 1; row < n; row++ {
		for col := 1; col < m; col++ {
			if matrix[row][col] != matrix[row-1][col-1] {
				return false
			}
		}
	}
	return true
}
