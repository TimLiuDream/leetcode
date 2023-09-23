package main

import "fmt"

func findCenter(edges [][]int) int {
	length := len(edges)
	m := make(map[int]int)
	for _, edge := range edges {
		for _, v := range edge {
			m[v]++
		}
	}
	for k, v := range m {
		if v == length {
			return k
		}
	}
	return -1
}

// 因为只有一个顶点会出现两次，所以这里直接判断相等即可
func findCenter1(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
	fmt.Println("----------------------")
	fmt.Println(findCenter1([][]int{{1, 2}, {2, 3}, {4, 2}}))
	fmt.Println(findCenter1([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}
