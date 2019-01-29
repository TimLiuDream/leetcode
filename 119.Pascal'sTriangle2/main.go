package main

import "fmt"

func getRow(rowIndex int) []int {
	numsGroup := make([][]int, 0)
	for i := 0; i <= rowIndex; i++ {
		if i == 0 {
			nums := []int{1}
			numsGroup = append(numsGroup, nums)
		} else if i == 1 {
			nums := []int{1, 1}
			numsGroup = append(numsGroup, nums)
		} else {
			nums := make([]int, i+1)
			nums[0] = 1
			nums[len(nums)-1] = 1
			preNums := numsGroup[i-1]
			for j := 1; j < len(nums)-1; j++ {
				nums[j] = preNums[j-1] + preNums[j]
			}
			numsGroup = append(numsGroup, nums)
		}
	}
	return numsGroup[rowIndex]
}

func main() {
	num := 3
	nums := getRow(num)
	fmt.Println(nums)
}
