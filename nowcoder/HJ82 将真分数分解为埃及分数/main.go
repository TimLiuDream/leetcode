package main

import (
	"fmt"
)

func main() {
	func2()
}

func func1() {
	var a, b int
	for {
		_, err := fmt.Scanf("%d/%d", &a, &b)
		if err != nil {
			break
		}
		for a != 1 {
			if a%b == 0 {
				fmt.Printf("%d/%d", 1, b/a)
				return
			}
			q := b / a
			fmt.Printf("%d/%d+", 1, 1+q)
			a = a - b%a
			b = b * (1 + q)
		}
		fmt.Printf("%d/%d\n", a, b)
	}
}

func func2() {
	var a, b int
	for {
		_, err := fmt.Scanf("%d/%d", &a, &b)
		if err != nil {
			break
		}
		for i := 0; i < a; i++ {
			if i == a-1 {
				fmt.Printf("%d/%d", 1, b)
				continue
			}
			fmt.Printf("%d/%d+", 1, b)
		}
	}
}
