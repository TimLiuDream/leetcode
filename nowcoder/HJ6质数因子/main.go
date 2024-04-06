package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			fmt.Printf("%d ", i)
		}
	}
	if n != 1 {
		fmt.Printf("%d ", n)
	}
}
