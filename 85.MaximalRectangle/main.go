package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(matrix))
}

// 单调栈实现
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	max := 0
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] = height[j] + 1
			}
		}
		max = int(math.Max(float64(max), float64(maxRectangle(height))))
	}
	return max
}

func maxRectangle(heights []int) int {
	maxarea := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxarea = int(math.Max(float64(maxarea), float64(heights[tmp]*(i-stack[len(stack)-1]-1))))
		}
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maxarea = int(math.Max(float64(maxarea), float64(heights[tmp]*(len(heights)-1-stack[len(stack)-1]))))
	}
	return maxarea
}
