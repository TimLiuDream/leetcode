package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := 0
	for i := 0; i <= n; i++ {
		if isZiShou1(i, i*i) {
			fmt.Println(i, i*i)
			ans++
		}
	}
	fmt.Println(ans)
}

func isZiShou(x, y int) bool {
	for x > 0 {
		t1 := x % 10
		t2 := y % 10
		if t1 != t2 {
			return false
		}
		x /= 10
		y /= 10
	}
	return true
}

func isZiShou1(x, y int) bool {
	return strings.HasSuffix(strconv.Itoa(y), strconv.Itoa(x))
}
