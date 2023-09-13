package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	for _, iNum := range nums1 {
		foundSelf := false
		part := -1
		for _, jNum := range nums2 {
			if iNum == jNum {
				foundSelf = true
				continue
			}
			if foundSelf && jNum > iNum {
				part = jNum
				break
			}
		}
		res = append(res, part)
	}
	return res
}

// 单调栈解法
func nextGreaterElement1(nums1, nums2 []int) []int {
	mp := make(map[int]int)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		// 构造单调栈
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[nums2[i]] = stack[len(stack)-1]
		} else {
			mp[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = mp[num]
	}
	return res
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
	fmt.Println(nextGreaterElement1(nums1, nums2))
	nums11 := []int{2, 4}
	nums21 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums11, nums21))
	fmt.Println(nextGreaterElement1(nums11, nums21))

	nums111 := []int{4}
	nums211 := []int{2, 5, 3, 6, 8, 4, 7, 1}
	fmt.Println(nextGreaterElement(nums111, nums211))
	fmt.Println(nextGreaterElement1(nums111, nums211))
}
