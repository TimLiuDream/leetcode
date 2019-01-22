package main

import "fmt"

func generate(numRows int) [][]int {
	numsGroup := make([][]int, numRows)
	for index := 0; index < numRows; index++ {
		if index == 0 {
			nums := []int{1}
			numsGroup[index] = nums
		} else if index == 1 {
			nums := []int{1, 1}
			numsGroup[index] = nums
		} else {
			nums := make([]int, index+1)
			nums[0] = 1
			nums[len(nums)-1] = 1
			for j := 1; j < len(nums)-1; j++ {
				nums[j] = numsGroup[index-1][j-1] + numsGroup[index-1][j]
			}
			numsGroup[index] = nums
		}
	}
	return numsGroup
}

func main() {
	num := 5
	numsGroup := generate(num)
	fmt.Println(numsGroup)
}
