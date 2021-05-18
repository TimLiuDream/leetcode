package main

import (
	"fmt"
)

// Label 位运算
// Label 数组
// Label 数学

func main() {
	arr := []int{2, 3, 1, 6, 7}
	fmt.Println(countTriplets(arr))
}

func countTriplets(arr []int) (ans int) {
	cnt := map[int]int{}
	total := map[int]int{}
	s := 0
	for k, val := range arr {
		if m, has := cnt[s^val]; has {
			ans += m*k - total[s^val]
		}
		cnt[s]++
		total[s] += k
		s ^= val
	}
	return
}
