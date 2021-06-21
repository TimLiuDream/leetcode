package main

import (
	"fmt"
	"math/bits"
)

// Label: 位运算
// Label: 回溯算法

func main() {
	turnedOn := 1
	fmt.Println(readBinaryWatch(turnedOn))
}

func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}
