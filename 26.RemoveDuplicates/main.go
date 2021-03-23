package main

import "fmt"

func removeDuplicates(nums []int) int {
	//如果是空切片，那就返回0
	if len(nums) == 0 {
		return 0
	}
	//用两个标记来比较相邻位置的值
	//当一样的话，那就不管继续
	//当不一样的时候，就把right指向的值赋值给left下一位
	left, right := 0, 1
	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	fmt.Println(nums[:left+1])
	return left + 1
}

//不考虑空间复杂度
//引入map，把不一样的放进去
func removeDuplicates1(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = v
		}
	}
	return len(m)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))
}
