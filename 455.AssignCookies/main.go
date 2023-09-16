package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))

	g1 := []int{1, 2}
	s1 := []int{1, 2, 3}
	fmt.Println(findContentChildren(g1, s1))

	g2 := []int{10, 9, 8, 7}
	s2 := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g2, s2))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	result := 0
	for i := 0; i < len(s) && len(g) > 0; i++ {
		if s[i] >= g[0] {
			g = g[1:]
			result++
		}
	}
	return result
}
