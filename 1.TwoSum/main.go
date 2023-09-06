package main

import "fmt"

// Label 数组
// Label 哈希表

func main() {
	slice := []int{2, 3, 4}
	a := twoSum1(slice, 6)
	fmt.Println(a)
}

// 双循环
func twoSum(slice []int, i int) []int {
	for iIndex := 0; iIndex < len(slice); iIndex++ {
		for jIndex := iIndex + 1; jIndex < len(slice); jIndex++ {
			if slice[iIndex]+slice[jIndex] == i {
				return []int{iIndex, jIndex}
			}
		}
	}
	return []int{}
}

// 将数字先存储到 map 中，然后再循环 slice，获取的差值在 map 中查找
func twoSum1(slice []int, i int) []int {
	m := make(map[int]int)
	for index, value := range slice {
		m[value] = index
	}
	for index, value := range slice {
		sub := i - value
		mIndex, ok := m[sub]
		if mIndex == index || !ok {
			continue
		}
		return []int{index, mIndex}
	}
	return []int{}
}

// 只循环一次，用 map 辅助记录，如果差值存在就返回，不存在就添加到 map 中
func twoSum2(slice []int, i int) []int {
	m := make(map[int]int)
	for index, value := range slice {
		mIndex, ok := m[i-value]
		if ok {
			return []int{mIndex, index}
		}
		m[value] = index
	}
	return []int{}
}
