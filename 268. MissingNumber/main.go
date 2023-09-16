package main

import (
	"fmt"
	"sort"
)

// 使用 map 记录解法
func missingNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

// 排序线性解法
func missingNumber1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)+1; i++ {
		if i != nums[i] {
			return i
		}
		if i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}

// 通过比较总和
func missingNumber2(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}

// 二分查找
func missingNumber3(nums []int) int {
	// 首先对数组进行排序
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		// 如果中间元素等于索引值，说明缺失的数字在右侧
		if nums[mid] == mid {
			left = mid + 1
		} else {
			// 否则，缺失的数字在左侧或者当前位置就是缺失的数字
			right = mid - 1
		}
	}

	// 返回 left 作为缺失的数字
	return left
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{0}))
	fmt.Println("=======================")
	fmt.Println(missingNumber1([]int{3, 0, 1}))
	fmt.Println(missingNumber1([]int{0, 1}))
	fmt.Println(missingNumber1([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber1([]int{0}))
	fmt.Println("=======================")
	fmt.Println(missingNumber2([]int{3, 0, 1}))
	fmt.Println(missingNumber2([]int{0, 1}))
	fmt.Println(missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber2([]int{0}))
	fmt.Println("=======================")
	fmt.Println(missingNumber3([]int{3, 0, 1}))
	fmt.Println(missingNumber3([]int{0, 1}))
	fmt.Println(missingNumber3([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber3([]int{0}))
}
