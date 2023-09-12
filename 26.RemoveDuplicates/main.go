package main

import "fmt"

// 双指针解法
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//用两个标记来比较相邻位置的值
	//当一样的话，那就不管继续
	//当不一样的时候，就把right指向的值赋值给left下一位
	left, right := 0, 0
	for right < len(nums)-1 {
		right++
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	return left + 1
}

// 使用 map 和计数器
func removeDuplicates1(nums []int) int {
	m := make(map[int]struct{})
	counter := 0 // 记录有多少个不一致的数字
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
			nums[counter] = num
			counter++
		}
	}
	return counter
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))
	nums2 := []int{1, 1}
	fmt.Println(removeDuplicates(nums2))
	nums3 := []int{}
	fmt.Println(removeDuplicates(nums3))
	fmt.Println("----------------------------")
	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates1(nums))
	nums1 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates1(nums1))
	nums2 = []int{1, 1}
	fmt.Println(removeDuplicates1(nums2))
	nums3 = []int{}
	fmt.Println(removeDuplicates1(nums3))
}
