package main

import "fmt"

func searchInsert(nums []int, target int) int {
	//当数组中最大的那个数都比target小的话，那就返回数组长度就好了
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	//进行for循环，当对应下标的数>=target的时候，那么就返回此时的下标
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			result = i
			break
		}
	}
	return result
}

func searchInsert1(nums []int, target int) int {
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) >> 1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				end = mid - 1
			}
		} else {
			start = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
	fmt.Println(searchInsert(nums, 0))

	fmt.Println(searchInsert1(nums, 5))
	fmt.Println(searchInsert1(nums, 2))
	fmt.Println(searchInsert1(nums, 7))
	fmt.Println(searchInsert1(nums, 0))
}
