package main

import (
	"fmt"
	"math"
)

func main() {
	var v float64
	fmt.Scan(&v)
	flag := 1
	if v < 0 {
		flag = -1
		v = math.Abs(v)
	}

	var (
		a float64 = 0
		b float64 = 3
	)
	for a <= b {
		mid := (b-a)/2 + a
		if mid*mid*mid >= v {
			if (mid-0.05)*(mid-0.05)*(mid-0.05) < v {
				fmt.Printf("%.1f", mid*float64(flag))
				return
			}
			b = mid
		} else {
			a = mid
		}
	}
}
