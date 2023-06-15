package main

import "fmt"

func main() {
	nums1, k1 := []int{1, 2, 3, 1}, 3
	fmt.Println(containsNearbyDuplicate(nums1, k1))

	nums2, k2 := []int{1, 0, 1, 1}, 1
	fmt.Println(containsNearbyDuplicate(nums2, k2))

	nums3, k3 := []int{1, 2, 3, 1, 2, 3}, 2
	fmt.Println(containsNearbyDuplicate(nums3, k3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		index, ok := m[num]
		if ok && abs(index, i) <= k {
			return true
		}
		m[num] = i
	}
	return false
}

func abs(a, b int) int {
	result := a - b
	if result < 0 {
		return -1 * result
	}
	return result
}
