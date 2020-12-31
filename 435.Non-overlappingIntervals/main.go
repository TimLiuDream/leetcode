package main

import (
	"fmt"
	"sort"
)

func main() {
	// intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	// fmt.Println(eraseOverlapIntervals(intervals))

	// intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
	// fmt.Println(eraseOverlapIntervals(intervals))

	intervals := [][]int{{1, 2}, {2, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	placeholder := make([]int, n)
	for i := range placeholder {
		placeholder[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				placeholder[i] = max(placeholder[i], placeholder[j]+1)
			}
		}
	}
	return n - max(placeholder...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
