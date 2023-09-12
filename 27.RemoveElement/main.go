package main

import "fmt"

func removeElement(nums []int, val int) int {
	//如果是空切片，那就返回0
	if len(nums) == 0 {
		return 0
	}
	//用一个索引
	//循环去比较
	//当一样的时候就删除对应下标的值
	//当不一样的时候，索引加1
	index := 0
	for index < len(nums) {
		if nums[index] == val {
			nums = append(nums[:index], nums[index+1:]...)
			continue
		}
		index++
	}
	return len(nums)
}

// 双指针
func removeElement1(nums []int, val int) int {
	left, right := 0, 0
	for ; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// 双指针优化，使用头尾指针
func removeElement2(nums []int, val int) int {
	head, tail := 0, len(nums)
	for head < tail {
		if nums[head] == val {
			nums[head] = nums[tail-1]
			tail--
		} else {
			head++
		}
	}
	return head
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement1(nums, 3))
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement1(nums1, 2))
	nums2 := []int{}
	fmt.Println(removeElement1(nums2, 2))
	fmt.Println("--------------------------")
	nums = []int{3, 2, 2, 3}
	fmt.Println(removeElement2(nums, 3))
	nums1 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement2(nums1, 2))
	nums2 = []int{1}
	fmt.Println(removeElement2(nums2, 1))
}
