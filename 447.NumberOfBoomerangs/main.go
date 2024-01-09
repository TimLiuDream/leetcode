package main

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	result := 0
	for _, point := range points {
		m := make(map[int]int)
		for _, q := range points {
			cnt := ((point[0] - q[0]) * (point[0] - q[0])) + ((point[1] - q[1]) * (point[1] - q[1]))
			m[cnt]++
		}
		for _, v := range m {
			result += v * (v - 1)
		}
	}
	return result
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}}))
}
