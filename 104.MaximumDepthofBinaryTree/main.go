package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := new(TreeNode)
	root.Val = 3
	n1 := new(TreeNode)
	n1.Val = 9
	n2 := new(TreeNode)
	n2.Val = 20
	n3 := new(TreeNode)
	n3.Val = 15
	n4 := new(TreeNode)
	n4.Val = 7

	root.Left = n1
	root.Right = n2
	n2.Left = n3
	n2.Right = n4

	fmt.Println(maxDepth(root))
}
