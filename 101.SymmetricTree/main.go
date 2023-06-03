package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return equal(root, root)
}

func equal(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && equal(t1.Left, t2.Right) && equal(t1.Right, t2.Left)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return equal1(root.Left, root.Right)
}

func equal1(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && equal1(t1.Left, t2.Right) && equal1(t1.Right, t2.Left)
	}
	return false
}

func main() {
	root := new(TreeNode)
	root.Val = 1

	l1l := new(TreeNode)
	l1l.Val = 2

	l1r := new(TreeNode)
	l1r.Val = 2

	l11r := new(TreeNode)
	l11r.Val = 3

	l21r := new(TreeNode)
	l21r.Val = 3

	root.Left = l1l
	root.Right = l1r
	l1l.Left = l11r
	l1r.Right = l21r

	fmt.Println(isSymmetric(root))
	fmt.Println(isSymmetric1(root))
}
