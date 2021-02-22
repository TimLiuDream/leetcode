package main

import "fmt"

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	fmt.Println(removeStones(stones))
}

func removeStones(stones [][]int) int {
	fa, rank := map[int]int{}, map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if _, has := fa[x]; !has {
			fa[x], rank[x] = x, 1
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
	}

	for _, p := range stones {
		union(p[0], p[1]+10001)
	}
	ans := len(stones)
	for x, fx := range fa {
		if x == fx {
			ans--
		}
	}
	return ans
}
