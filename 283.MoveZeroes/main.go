package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums1 := []int{0}
	moveZeroes(nums1)
	fmt.Println(nums1)
	fmt.Println("=====")
	nums2 := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums2)
	fmt.Println(nums2)
	nums3 := []int{0}
	moveZeroes1(nums3)
	fmt.Println(nums3)
}

func moveZeroes(nums []int) {
	i, length := 0, len(nums)
	for i < length {
		if nums[i] == 0 {
			copy(nums, append(append(nums[:i], nums[i+1:]...), 0))
			length--
		} else {
			i++
		}
	}
}

// 双指针,除0以外的数往前移动，相对顺序不变
func moveZeroes1(nums []int) {
	i, j, length := 0, 0, len(nums)
	for j < length {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < length {
		nums[i] = 0
		i++
	}
}
