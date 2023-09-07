package main

import (
	"fmt"
	"math"
)

// 假设 t 分钟内可以将所有汽车都修理完，那么大于等于 t 分钟内都可以将所有汽车修理完。
// 假设 t 分钟内不能够将所有汽车都修理完，那么小于等于 t 分钟内也不能够将所有汽车修理完。
func repairCars(ranks []int, cars int) int64 {
	l, r := 0, math.MaxInt64

	var check = func(v int) bool {
		cnt := 0
		for _, x := range ranks {
			// 原公式为： ranks[i] * carNum * carNum
			// cnt 为个人能修理汽车的数量
			cnt += int(math.Sqrt(float64(v / x)))
		}
		return cnt >= cars
	}

	for l < r {
		v := (l + r) / 2 // 取中点
		if check(v) {
			r = v
		} else {
			l = v + 1
		}
	}

	return int64(l)
}

func main() {
	fmt.Println(repairCars([]int{4, 2, 3, 1}, 10))
	fmt.Println(repairCars([]int{5, 1, 8}, 6))
}
