package main

import "fmt"

// 一次遍历
func search(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// 二分查找
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			if nums[mid] == target {
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println("--------------------")
	fmt.Println(search1([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search1([]int{-1, 0, 3, 5, 9, 12}, 2))
}
