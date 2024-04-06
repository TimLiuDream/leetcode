package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	row := []int{1}
	for i := 1; i < n; i++ {
		row = append(row, row[i-1]+1+i)
	}

	m := [][]int{row}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if i != 0 {
				m = append(m, make([]int, n-i))
				m[i][j] = m[i-1][j+1] - 1
			}
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}
