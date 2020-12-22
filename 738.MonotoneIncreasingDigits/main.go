package main

import (
	"fmt"
	"strconv"
)

func main() {
	N := 10
	// N := 1234
	// N := 332
	// N := 0
	// N := math.Pow10(9)
	fmt.Println(monotoneIncreasingDigits(N))
}

func monotoneIncreasingDigits(N int) int {
	s := []byte(strconv.Itoa(N))
	if len(s) <= 1 {
		return N
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[i+1] {
			s[i] -= 1
			for j := i + 1; j < len(s); j++ {
				s[j] = '9'
			}
		}
	}

	v, _ := strconv.Atoi(string(s))
	return v
}
