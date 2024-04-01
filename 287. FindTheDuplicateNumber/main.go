package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		count := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}

func findDuplicate1(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
	nums = []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
	nums = []int{3, 3, 3, 3, 3}
	fmt.Println(findDuplicate(nums))
}
