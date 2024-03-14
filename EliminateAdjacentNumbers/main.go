package main

import "fmt"

func eaNumber(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}
	stack := []int{}
	for _, num := range arr {
		if len(stack) == 0 {
			stack = append(stack, num)
			continue
		}
		if num+stack[len(stack)-1] == 10 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, num)
		}
	}
	return len(stack)
}

func buildChunks(arr []string, parts int) [][]string {
	count := len(arr) / parts
	chunks := make([][]string, 0)
	for i := 0; i <= parts; i++ {
		end := (i + 1) * count
		if end > len(arr) {
			end = len(arr)
		}
		chunk := arr[i*count : end]
		chunks = append(chunks, chunk)
	}
	return chunks
}

func main() {
	arr := []int{2, 1, 3, 7, 9, 2}
	fmt.Println(eaNumber(arr))
}
