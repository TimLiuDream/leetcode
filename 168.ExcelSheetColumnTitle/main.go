package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(703))
}

func convertToTitle(n int) string {
	if n <= 26 {
		return string(64 + n)
	}
	y := n % 26
	if y == 0 {
		return convertToTitle((n-y-1)/26) + convertToTitle(26)
	}
	return convertToTitle((n-y)/26) + convertToTitle(y)
}
