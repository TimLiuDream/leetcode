package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func main() {
	root := new(TreeNode)
	root.Val = 5
	l1 := new(TreeNode)
	l1.Val = 4
	l2 := new(TreeNode)
	l2.Val = 8
	l3 := new(TreeNode)
	l3.Val = 11
	l4 := new(TreeNode)
	l4.Val = 13
	l5 := new(TreeNode)
	l5.Val = 4
	l6 := new(TreeNode)
	l6.Val = 7
	l7 := new(TreeNode)
	l7.Val = 2
	l8 := new(TreeNode)
	l8.Val = 1

	root.Left = l1
	root.Right = l2
	l1.Left = l3
	l2.Left = l4
	l2.Right = l5
	l3.Left = l6
	l3.Right = l7
	l5.Right = l8

	fmt.Println(hasPathSum(root, 22))
}
