package main

import (
	"fmt"
)

func calculateArea(N, E int, instructions [][]int) int {
	totalArea := 0
	prevX, prevY, prevOffsetY := 0, 0, 0

	for i := 0; i < N; i++ {
		x, offsetY := instructions[i][0], instructions[i][1]

		y := prevY + prevOffsetY
		width := x - prevX
		area := cal(width, y)
		totalArea += area

		prevX = x
		prevY = y
		prevOffsetY = offsetY
	}

	lastWidth := E - prevX
	lastY := prevY + prevOffsetY
	lastArea := cal(lastWidth, lastY)
	totalArea += lastArea

	return totalArea
}

func cal(width, y int) int {
	if width < 0 {
		width *= -1
	}
	if y < 0 {
		y *= -1
	}
	return width * y
}

func main() {
	// 输入示例
	N, E := 4, 10
	instructions := [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, -2},
	}

	// 计算并输出面积
	area := calculateArea(N, E, instructions)
	fmt.Println("The calculated area is:", area)

	N, E = 2, 4
	instructions = [][]int{
		{0, 1},
		{2, -2},
	}
	area = calculateArea(N, E, instructions)
	fmt.Println("The calculated area is:", area)
}
