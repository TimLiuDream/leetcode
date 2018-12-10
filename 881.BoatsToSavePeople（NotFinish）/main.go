package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	//排序
	sort.Ints(people)

	a, b, count := 0, len(people)-1, 0
	for a <= b {
		if people[b]+people[a] <= limit {
			a++
		}
		b--
		count++
	}
	return count
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
	fmt.Println(numRescueBoats([]int{5, 1, 4, 2}, 6))
}
