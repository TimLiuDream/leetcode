package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a, b := n/2, n/2
	for a >= 0 && b <= n {
		if isValid(a) && isValid(b) {
			fmt.Println(a)
			fmt.Println(b)
			return
		}
		a--
		b++
	}
}

func isValid(num1 int) bool {
	num := 0
	for i := 1; i <= num1; i++ {
		if num1%i == 0 {
			num++
		}
	}

	if num == 2 {
		return true
	}

	return false
}
