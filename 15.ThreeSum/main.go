package main

import (
	"fmt"
	"sort"
)

//此方法未完成
func threeSum1(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	m := make(map[int]int)
	//先把源数组保存进map中
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			v, ok := m[-(nums[i] + nums[j])]
			if ok {
				//为了去重，删除掉nums[i],nums[j],-(nums[i]+nums[j])的key
				delete(m, nums[i])
				delete(m, nums[j])
				delete(m, -(nums[i] + nums[j]))
				item := []int{nums[i], nums[j], v}
				result = append(result, item)
				continue
			}
		}
	}
	return result
}

func threeSum2(nums []int) [][]int {
	//先对数组排序
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		z := len(nums) - 1
		for z > j {
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}
				j++
				z--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum1(nums))
}
