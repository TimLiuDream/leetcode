package main

import (
	"fmt"
	"sort"
)

// 分数互不相同
// 排序解法
func findRelativeRanks(score []int) []string {
	// 将分数以及对应的索引转换成 map
	m := map[int]int{}
	for index, s := range score {
		m[s] = index
	}
	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	result := make([]string, len(score), len(score))
	for rank, s := range score {
		title := title(rank + 1)
		index := m[s]
		result[index] = title
	}
	return result
}

func title(rank int) (result string) {
	if rank == 1 {
		result = "Gold Medal"
	} else if rank == 2 {
		result = "Silver Medal"
	} else if rank == 3 {
		result = "Bronze Medal"
	} else {
		result = fmt.Sprintf("%d", rank)
	}
	return result
}

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
