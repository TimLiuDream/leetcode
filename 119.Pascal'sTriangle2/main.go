package main

import "fmt"

func getRow(rowIndex int) []int {
	numsGroup := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		if i == 0 {
			numsGroup[i] = []int{1}
		} else if i == 1 {
			numsGroup[i] = []int{1, 1}
		} else {
			nums := make([]int, i+1)
			nums[0], nums[len(nums)-1] = 1, 1
			for j := 1; j < len(nums)-1; j++ {
				nums[j] = numsGroup[i-1][j-1] + numsGroup[i-1][j]
			}
			numsGroup[i] = nums
		}
	}
	return numsGroup[rowIndex]
}

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}
