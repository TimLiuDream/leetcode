package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2] //因为题目要求众数是说元素的个数大于n/2的
}

func majorityElement1(nums []int) int {
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}

	maxCount := 0
	maxKey := 0
	for key, value := range numCountMap {
		if value >= maxCount {
			maxKey = key
			maxCount = value
		}
	}
	return maxKey
}

func majorityElement2(nums []int) int {
	result, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			result = num
			count++
		} else if result != num {
			count--
		} else {
			count++
		}
	}
	return result
}

// 分治思想：如果 a 是数组 nums 的众数，那么 a 必然是 nums 分成两半后中一半的众数
func majorityElement3(nums []int) int {
	return majorityElementR(nums, 0, len(nums)-1)
}

func majorityElementR(nums []int, leftIndex, rightIndex int) int {
	if leftIndex == rightIndex {
		return nums[leftIndex]
	}
	mid := (rightIndex-leftIndex)/2 + leftIndex
	leftValue := majorityElementR(nums, leftIndex, mid)
	rightValue := majorityElementR(nums, mid+1, rightIndex)
	if leftValue == rightValue {
		return leftValue
	}
	lc, rc := count(nums, leftValue, leftIndex, rightIndex), count(nums, rightValue, leftIndex, rightIndex)
	if lc > rc {
		return leftValue
	}
	return rightValue
}

func count(nums []int, num, leftIndex, rightIndex int) int {
	count := 0
	for i := leftIndex; i <= rightIndex; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}

func main() {
	nums := []int{3, 2, 3}
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
	fmt.Println(majorityElement(nums1))

	fmt.Println(majorityElement1(nums))
	fmt.Println(majorityElement1(nums1))

	fmt.Println(majorityElement2(nums))
	fmt.Println(majorityElement2(nums1))

	fmt.Println(majorityElement3(nums))
	fmt.Println(majorityElement3(nums1))
}
