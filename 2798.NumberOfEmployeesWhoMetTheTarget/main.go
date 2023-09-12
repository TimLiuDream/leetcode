package main

import (
	"fmt"
	"sort"
)

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	res := 0
	for _, hour := range hours {
		if hour >= target {
			res++
		}
	}
	return res
}

func numberOfEmployeesWhoMetTarget1(hours []int, target int) int {
	sort.Ints(hours)
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			return len(hours) - i
		}
	}
	return 0
}

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{5, 1, 4, 2, 2}, 6))
	fmt.Println(numberOfEmployeesWhoMetTarget1([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(numberOfEmployeesWhoMetTarget1([]int{5, 1, 4, 2, 2}, 6))
}
