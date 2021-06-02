package main

import "fmt"

// Label 数学
// Label 动态规划
// Label 前缀和

// solution https://leetcode-cn.com/problems/continuous-subarray-sum/solution/lian-xu-de-zi-shu-zu-he-by-leetcode-solu-rdzi/

func main() {
	nums := []int{23, 2, 4, 6, 7}
	k := 6
	fmt.Println(checkSubarraySum(nums, k))

	nums = []int{23, 2, 6, 4, 7}
	k = 6
	fmt.Println(checkSubarraySum(nums, k))

	nums = []int{23, 2, 6, 4, 7}
	k = 13
	fmt.Println(checkSubarraySum(nums, k))
}

func checkSubarraySum(nums []int, k int) bool {
	m := len(nums)
	if m < 2 {
		return false
	}
	mp := map[int]int{0: -1}
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if prevIndex, has := mp[remainder]; has {
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}
	return false
}
