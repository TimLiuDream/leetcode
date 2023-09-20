package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		res := make([]int, 0)
		if node == nil {
			return res
		}
		res = append(res, dfs(node.Left)...)
		res = append(res, node.Val)
		res = append(res, dfs(node.Right)...)
		return res
	}
	result := dfs(root)
	min := math.MaxInt64
	for i := 1; i < len(result); i++ {
		if v := abs(result[i] - result[i-1]); v < min {
			min = v
		}
	}
	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(getMinimumDifference(root))

	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}
	fmt.Println(getMinimumDifference(root1))
}
