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

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums1, 2))
}
