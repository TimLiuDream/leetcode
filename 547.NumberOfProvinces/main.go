package main

import "fmt"

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}

func findCircleNum(isConnected [][]int) int {
	vins := make([]bool, len(isConnected))
	var dfs func(from int)
	dfs = func(from int) {
		vins[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vins[to] {
				dfs(to)
			}
		}
	}
	ans := 0
	for i, v := range vins {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans
}
