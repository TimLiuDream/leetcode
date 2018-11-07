package main

import "fmt"

//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[target-nums[i]]
		if ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}
	return []int{0}
}

func main() {
	slice := []int{2, 3, 4}
	a := TwoSum(slice, 6)
	fmt.Println(a)
}
