package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lBalanced := isBalanced(root.Left)
	rBalanced := isBalanced(root.Right)
	sub := abs(maxDepth(root.Left) - maxDepth(root.Right))
	return sub <= 1 && lBalanced && rBalanced
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lDepth := maxDepth(node.Left)
	rDepth := maxDepth(node.Right)
	return 1 + max(lDepth, rDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(isBalanced(nil))

	root1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 20},
		},
	}
	fmt.Println(isBalanced(root1))

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isBalanced(root2))
}
