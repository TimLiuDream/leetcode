package main

import "fmt"

func main() {
	// nums := []int{1, 3}
	// n := 6
	// fmt.Println(minPatches(nums, n))

	nums := []int{1, 5, 10}
	n := 20
	fmt.Println(minPatches(nums, n))

	// nums := []int{1,2,2}
	// n := 5
	// fmt.Println(minPatches(nums, n))
}

func minPatches(nums []int, n int) int {
	patches := 0
	i := 0
	miss := 1
	l := len(nums)
	for miss <= n {
		if i < l && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			patches++
		}
	}
	return patches
}
