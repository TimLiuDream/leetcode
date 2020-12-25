package main

import (
	"fmt"
	"sort"
)

func main() {
	// g := []int{1, 2, 3}
	// s := []int{1, 1}
	// fmt.Println(findContentChildren(g, s))

	// g := []int{1, 2}
	// s := []int{1, 2, 3}
	// fmt.Println(findContentChildren(g, s))

	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	result := 0
	i := 0
	for _, sValue := range s {
		for i <= len(g)-1 {
			if sValue >= g[i] {
				result++
				i++
			}
			break
		}
	}
	return result
}
