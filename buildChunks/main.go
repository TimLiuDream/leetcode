package main

import "fmt"

// 均分数组为k份 然后相对顺序不变 每份尽可能平均

func buildChunks(arr []int, k int) [][]int {
	length := len(arr)
	count := 0
	if length%k == 0 {
		count = length / k
	} else {
		count = (length / k) + 1
	}
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		end := (i + 1) * count
		if end > length {
			end = length
		}
		result[i] = arr[i*count : end]
	}
	return result
}

func main() {
	arr := []int{2, 1, 3, 7, 9, 2}
	fmt.Println(buildChunks(arr, 4))
}
