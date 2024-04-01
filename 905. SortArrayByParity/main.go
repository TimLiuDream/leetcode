package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	a, b := []int{}, []int{}
	for _, num := range nums {
		if num%2 == 1 {
			b = append(b, num)
		} else {
			a = append(a, num)
		}
	}
	return append(a, b...)
}

func sortArrayByParity1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 1 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%2 != 0 {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}
	return nums
}

func sortArrayByParity2(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for _, num := range nums {
		if num%2 == 0 {
			ans = append(ans, num)
		}
	}
	for _, num := range nums {
		if num%2 == 1 {
			ans = append(ans, num)
		}
	}
	return ans
}

// 双指针 + 一次遍历
func sortArrayByParity3(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 0, n-1
	for _, num := range nums {
		if num%2 == 0 {
			ans[left] = num
			left++
		} else {
			ans[right] = num
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity1([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
	fmt.Println(sortArrayByParity1([]int{0}))
}
