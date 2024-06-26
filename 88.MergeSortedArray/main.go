package main

import (
	"fmt"
	"sort"
)

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)

	tIndex, jIndex := 0, 0
	for i := 0; i < len(nums1); i++ {
		if tIndex >= len(temp) {
			nums1[i] = nums2[jIndex]
			jIndex++
			continue
		}
		if jIndex >= n {
			nums1[i] = temp[tIndex]
			tIndex++
			continue
		}
		if temp[tIndex] <= nums2[jIndex] {
			nums1[i] = temp[tIndex]
			tIndex++
		} else {
			nums1[i] = nums2[jIndex]
			jIndex++
		}
	}
}

// 合并后排序
func merge1(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}

// 直接合并后排序
func merge3(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 普通双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := 0, 0
	result := make([]int, 0, m+n)
	for {
		if p1 == m {
			result = append(result, nums2[p2:]...)
			break
		}
		if p2 == n {
			result = append(result, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			result = append(result, nums1[p1])
			p1++
		} else {
			result = append(result, nums2[p2])
			p2++
		}
	}
	copy(nums1, result)
}

func merge4(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
	fmt.Println(num1)

	num1 = []int{1, 2, 3, 0, 0, 0}
	num2 = []int{2, 5, 6}
	merge1(num1, 3, num2, 3)
	fmt.Println(num1)

	num1 = []int{1, 2, 3, 0, 0, 0}
	num2 = []int{2, 5, 6}
	merge2(num1, 3, num2, 3)
	fmt.Println(num1)
}
