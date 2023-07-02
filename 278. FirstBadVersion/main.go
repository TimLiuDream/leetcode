package main

import (
	"fmt"
	"sort"
)

func main() {
	i := (0 + 1) / 2
	fmt.Println(i)
	fmt.Println(firstBadVersion(5))
	fmt.Println(firstBadVersion(1))
}

func firstBadVersion(n int) int {
	l, r := 0, n
	for l <= r {
		mid := (l + r) / 2
		ok := isBadVersion(mid)
		if ok {
			if !isBadVersion(mid - 1) {
				return mid
			} else {
				r = mid - 1
			}
		} else {
			l = mid + 1
		}
	}
	return -1
}

func firstBadVersion1(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
