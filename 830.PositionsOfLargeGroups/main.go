package main

import "fmt"

func main() {
	// s := "abbxxxxzzy"
	// fmt.Println(largeGroupPositions(s))

	// s := "abc"
	// fmt.Println(largeGroupPositions(s))

	// s := "abcdddeeeeaabbbcd"
	// fmt.Println(largeGroupPositions(s))

	// s := "aba"
	// fmt.Println(largeGroupPositions(s))

	s := "aaa"
	fmt.Println(largeGroupPositions(s))
}

func largeGroupPositions(s string) [][]int {
	result := make([][]int, 0)
	if len(s) == 0 {
		return result
	}
	stack := make([]rune, 0)
	for i, charster := range s {
		if len(stack) == 0 {
			stack = append(stack, charster)
		} else {
			if charster == stack[0] {
				stack = append(stack, charster)
			} else {
				if len(stack) >= 3 {
					result = append(result, []int{i - len(stack), i - 1})
				}
				stack = make([]rune, 0)
				stack = append(stack, charster)
			}
		}
	}
	if len(stack) >= 3 {
		result = append(result, []int{len(s) - len(stack), len(s) - 1})
	}
	return result
}
