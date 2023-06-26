package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 || n == 2 {
		return true
	}
	if n%2 != 0 {
		return false
	}
	return isPowerOfTwo(n / 2)
}
