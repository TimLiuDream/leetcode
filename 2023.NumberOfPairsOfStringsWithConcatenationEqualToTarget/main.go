package main

import "fmt"

func numOfPairs(nums []string, target string) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				result++
			}
		}
	}
	return result
}

func numOfPairs1(nums []string, target string) int {
	ans := 0
	cnt := map[string]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for i, n := 1, len(target); i < n; i++ {
		p, s := target[:i], target[i:] // 枚举所有前缀+后缀
		if p != s {
			ans += cnt[p] * cnt[s]
		} else {
			ans += cnt[p] * (cnt[p] - 1) // 前后缀相同时，相当于从 cnt[p] 个下标中选择两个不同下标的排列数，即 A(cnt[p], 2)
		}
	}
	return ans
}

func main() {
	fmt.Println(numOfPairs([]string{"777", "7", "77", "77"}, "7777"))
	fmt.Println(numOfPairs([]string{"123", "4", "12", "34"}, "1234"))
	fmt.Println(numOfPairs([]string{"1", "1", "1"}, "11"))
}
