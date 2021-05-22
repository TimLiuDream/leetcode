package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(xorGame(nums))
}

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}
