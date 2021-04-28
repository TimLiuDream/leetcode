package main

import (
	"fmt"
	"math"
)

// Label 数学

func main() {
	c := 5
	fmt.Println(judgeSquareSum(c))
	c = 3
	fmt.Println(judgeSquareSum(c))
	c = 4
	fmt.Println(judgeSquareSum(c))
	c = 2
	fmt.Println(judgeSquareSum(c))
	c = 1
	fmt.Println(judgeSquareSum(c))
}

func judgeSquareSum(c int) bool {
	for i := 0; i < c; i++ {
		v := math.Sqrt(float64(c - i*i))
		if v == math.Floor(v){
			return true
		}
	}
	return false
}
