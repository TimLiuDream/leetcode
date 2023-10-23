package main

import "fmt"

func rotate(nums []int, k int) {
	// 处理边界情况
	if len(nums) == 0 || k <= 0 {
		return
	}

	// 创建临时数组并复制元素
	tmp := make([]int, len(nums))
	copy(tmp, nums)

	// 将临时数组的元素从右往左依次复制回原始数组
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[(i-k+len(nums))%len(nums)]
	}
}

// 数组右移
func main() {
	// 测试示例
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println("原始数组:", nums)
	rotate(nums, k)
	fmt.Println("右移后的数组:", nums)
}
