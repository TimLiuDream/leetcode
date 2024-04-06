package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := 1
	if n == 1 {
		fmt.Println(a)
		return
	}
	for i := 2; i <= n; i++ {
		tmp := ""
		for j := 0; j < i; j++ {
			a += 2
			tmp += fmt.Sprintf("%d+", a)
		}
		if i == n {
			fmt.Println(strings.TrimRight(tmp, "+"))
		}
	}
}
