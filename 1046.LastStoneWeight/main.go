package main

import (
	"fmt"
	"sort"
)

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))

	// stones := []int{2, 2}
	// fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		v := stones[0] - stones[1]
		if v != 0 {
			stones = append(stones[2:], v)
		} else {
			if len(stones) == 2 {
				return v
			} else {
				stones = stones[2:]
			}
		}
	}
	return stones[0]
}
