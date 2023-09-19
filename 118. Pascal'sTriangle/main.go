package main

import "fmt"

func generate(numRows int) [][]int {
	numsGroup := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			numsGroup[0] = []int{1}
		} else if i == 1 {
			numsGroup[1] = []int{1, 1}
		} else {
			numsGroup[i] = make([]int, i+1)
			numsGroup[i][0], numsGroup[i][len(numsGroup[i])-1] = 1, 1
			for j := 1; j < i; j++ {
				numsGroup[i][j] = numsGroup[i-1][j-1] + numsGroup[i-1][j]
			}
		}
	}
	return numsGroup
}

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(5))
}
