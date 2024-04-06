package main

import (
	"fmt"
)

func main() {
	for {
		n := 0
		_, err := fmt.Scan(&n)
		n1 := 0
		if err != nil {
			break
		}
		for i := 1; i <= n; i++ {
			li := make([]int, 0)
			sum := 0
			for j := 1; j < i; j++ {
				if i%j == 0 {
					li = append(li, j)
				}
			}
			for _, k := range li {
				sum += k
			}
			if sum == i {
				n1++
			}
		}
		fmt.Println(n1)
	}
}
