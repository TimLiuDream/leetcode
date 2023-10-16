package main

import "fmt"

func distributeCandies(candyType []int) int {
	m := make(map[int]struct{})
	for _, ct := range candyType {
		m[ct] = struct{}{}
	}
	result := make([]int, 0)
	for ct, _ := range m {
		if len(result) < len(candyType)/2 {
			result = append(result, ct)
		}
	}
	return len(result)
}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(distributeCandies([]int{6, 6, 6, 6}))
}
