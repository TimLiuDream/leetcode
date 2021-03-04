package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	// i := 0
	// j := 2
	// obj := Constructor(nums)
	// p := obj.SumRange(i, j)
	// fmt.Println(p)

	// i := 2
	// j := 5
	// obj := Constructor(nums)
	// p := obj.SumRange(i, j)
	// fmt.Println(p)

	i := 0
	j := 5
	obj := Constructor(nums)
	p := obj.SumRange(i, j)
	fmt.Println(p)
}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

func (na *NumArray) SumRange(i, j int) int {
	return na.sums[j+1] - na.sums[i]
}
