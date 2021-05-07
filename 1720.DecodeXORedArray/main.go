package main

import "fmt"

// Label 位运算

func main() {
	encoded := []int{1, 2, 3}
	first := 1
	fmt.Println(decode(encoded, first))
	encoded = []int{6, 2, 7, 3}
	first = 4
	fmt.Println(decode(encoded, first))
}

func decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)
	result[0] = first
	for i, encode := range encoded {
		result[i+1] = encode ^ result[i]
	}
	return result
}
