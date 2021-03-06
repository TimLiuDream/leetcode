package main

import (
	"fmt"
)

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	//把nums1复制到temp中
	temp := make([]int, m)
	copy(temp, nums1)

	tIndex, nums2Indes := 0, 0 //t为temp的索引，j为nums2的索引
	for i := 0; i < len(nums1); i++ {
		//当tIndex大于temp的长度，那就是说temp全部放进去了nums1中，那剩下的就是放nums2剩余的值了
		if tIndex >= len(temp) {
			nums1[i] = nums2[nums2Indes]
			nums2Indes++
			continue
		}
		//当nums2Indes大于nums2的长度的时候，那就是说明nums2全部都放进去了nums1中，那剩下的就是放temp剩余的值了
		if nums2Indes >= n {
			nums1[i] = temp[tIndex]
			tIndex++
			continue
		}
		//比较nums2与temp对应值的大小，小的那个就放进nums1中
		if nums2[nums2Indes] <= temp[tIndex] {
			nums1[i] = nums2[nums2Indes]
			nums2Indes++
		} else {
			nums1[i] = temp[tIndex]
			tIndex++
		}
	}
	fmt.Println(nums1)
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
}
