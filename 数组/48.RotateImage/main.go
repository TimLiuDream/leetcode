package main

import "fmt"

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for rowIndex := 0; rowIndex < n/2; rowIndex++ {
		for columnIndex := 0; columnIndex < (n+1)/2; columnIndex++ {
			matrix[rowIndex][columnIndex],
				matrix[n-columnIndex-1][rowIndex],
				matrix[n-rowIndex-1][n-columnIndex-1],
				matrix[columnIndex][n-rowIndex-1] =
				matrix[n-columnIndex-1][rowIndex],
				matrix[n-rowIndex-1][n-columnIndex-1],
				matrix[columnIndex][n-rowIndex-1],
				matrix[rowIndex][columnIndex]
		}
	}
}
