package main

import "fmt"

func giveGem(gem []int, operations [][]int) int {
	for _, operation := range operations {
		x := gem[operation[0]] / 2
		gem[operation[0]] = gem[operation[0]] - x
		gem[operation[1]] = gem[operation[1]] + x
	}
	max, min := gem[0], gem[0]
	for _, g := range gem {
		if g > max {
			max = g
		}
		if g < min {
			min = g
		}
	}
	return max - min
}

func main() {
	gem := []int{3, 1, 2}
	operations := [][]int{{0, 2}, {2, 1}, {2, 0}}
	fmt.Println(giveGem(gem, operations))
	gem1 := []int{100, 0, 50, 100}
	operations1 := [][]int{{0, 2}, {0, 1}, {3, 0}, {3, 0}}
	fmt.Println(giveGem(gem1, operations1))
	gem2 := []int{0, 0, 0, 0}
	operations2 := [][]int{{1, 2}, {3, 1}, {1, 2}}
	fmt.Println(giveGem(gem2, operations2))
}
