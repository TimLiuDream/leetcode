package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dep func(node *TreeNode) int
	dep = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lDepth := dep(node.Left)
		rDepth := dep(node.Right)
		ans = max(ans, lDepth+rDepth)
		return max(lDepth, rDepth) + 1
	}
	dep(root)
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
