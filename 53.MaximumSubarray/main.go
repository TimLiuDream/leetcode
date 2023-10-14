package main

import (
	"fmt"
)

// 动态规划
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 动态规划解法
// 假设sum[i]为以第i个元素结尾且和最大的连续子数组。
// 假设对于元素i，所有以它前面的元素结尾的子数组的长度都已经求得
// 那么以第i个元素结尾且和最大的连续子数组实际上，要么是以第i-1个元素结尾且和最大的连续子数组加上这个元素，要么是只包含第i个元素
// 即sum[i] = max(sum[j:i-1] + a[i], a[i])。//j是数组的某个下标，0<=j<i-1<n
// 可以通过判断sum[i-1] + a[i]是否大于a[i]来做选择，而这实际上等价于判断sum[i-1]是否大于0。
// 由于每次运算只需要前一次的结果，因此并不需要像普通的动态规划那样保留之前所有的计算结果，只需要保留上一次的即可，因此算法的时间和空间复杂度都很小
func maxSubArrayKMP(nums []int) int {
	res, sum := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		//这里是核心判断
		//当前几项大于0，那就相加
		if res > 0 {
			res += nums[i]
		} else { //如果不是，那就直接是当前项
			res = nums[i]
		}
		if sum < res {
			sum = res
		}
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
	nums1 := []int{-2, 1}
	fmt.Println(maxSubArray(nums1))
	nums2 := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums2))
}
