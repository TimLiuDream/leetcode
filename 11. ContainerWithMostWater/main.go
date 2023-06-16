package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxA := area(height, left, right)
	for left < right {
		maxA = max(maxA, area(height, left, right))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxA
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func area(height []int, i, j int) int {
	h := height[i]
	if height[i] > height[j] {
		h = height[j]
	}
	return h * (j - i)
}
