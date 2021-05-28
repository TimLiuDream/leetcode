package main

import (
	"fmt"
	"math/bits"
)

// Label 位运算

// solution https://leetcode-cn.com/problems/total-hamming-distance/solution/yi-ming-ju-chi-zong-he-by-leetcode-solut-t0ev/

func main() {
	nums := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(nums))

	nums = []int{4, 14, 4}
	fmt.Println(totalHammingDistance(nums))
}

func hammingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func totalHammingDistance(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			total += hammingDistance(nums[i], nums[j])
		}
	}
	return total
}
