package main

import (
	"fmt"
	"math"
)

// Label 数学

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = sign * x
	}
	result := 0
	for x > 0 {
		tmp := x % 10
		result = result*10 + tmp
		x = x / 10
	}
	if sign <= 0 {
		result = result * sign
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func reverse1(x int) (rev int) {
	for x != 0 {
		if x > math.MaxInt32/10 || x < math.MinInt32/10 {
			return
		}
		tmp := x % 10
		x /= 10
		rev = rev*10 + tmp
	}
	return
}

func main() {
	fmt.Println(reverse(-123123457687980))
}
