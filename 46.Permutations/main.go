package main

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	visited := make([]bool, len(nums))
	curr := make([]int, 0)
	backtracing(nums, visited, curr, &result)
	return result
}

func backtracing(nums []int, visited []bool, curr []int, result *[][]int) {
	if len(curr) == len(nums) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
		return
	}
	for i, num := range nums {
		if visited[i] {
			continue
		}
		visited[i] = true
		curr = append(curr, num)
		backtracing(nums, visited, curr, result)
		curr = curr[:len(curr)-1]
		visited[i] = false
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
