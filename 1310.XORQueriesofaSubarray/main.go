package main

import "fmt"

// Label 位运算符

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	fmt.Println(xorQueries1(arr, queries))

	arr = []int{4, 8, 2, 10}
	queries = [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
	fmt.Println(xorQueries1(arr, queries))
}

func xorQueries(arr []int, queries [][]int) []int {
	result := make([]int, 0)
	for _, query := range queries {
		v := 0
		for i, value := range arr {
			if i >= query[0] && i <= query[1] {
				v = v ^ value
			}
		}
		result = append(result, v)
	}
	return result
}

func xorQueries1(arr []int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i, query := range queries {
		v := 0
		a := arr[query[0] : query[1]+1]
		for _, value := range a {
			v ^= value
		}
		result[i] = v
	}
	return result
}
