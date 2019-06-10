package main

import (
	"fmt"

	"github.com/timliudream/leetcode"
)

func isSymmetric(root *leetcode.TreeNode) bool {
	return equal(root, root)
}

func equal(t1 *leetcode.TreeNode, t2 *leetcode.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && equal(t1.Left, t2.Right) && equal(t1.Right, t2.Left)
}

func main() {
	root := new(leetcode.TreeNode)
	root.Val = 1

	l1l := new(leetcode.TreeNode)
	l1l.Val = 2

	l1r := new(leetcode.TreeNode)
	l1r.Val = 2

	l11r := new(leetcode.TreeNode)
	l11r.Val = 3

	l21r := new(leetcode.TreeNode)
	l21r.Val = 3

	root.Left = l1l
	root.Right = l1r
	l1l.Left = l11r
	l1r.Right = l21r

	fmt.Println(isSymmetric(root))
}
