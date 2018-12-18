package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

func main() {
	root := TreeNode{}
	root.Val = 6
	node1 := TreeNode{}
	node1.Val = 2
	node2 := TreeNode{}
	node2.Val = 8
	node3 := TreeNode{}
	node3.Val = 0
	node4 := TreeNode{}
	node4.Val = 4
	node5 := TreeNode{}
	node5.Val = 7
	node6 := TreeNode{}
	node6.Val = 9
	node7 := TreeNode{}
	node7.Val = 3
	node8 := TreeNode{}
	node8.Val = 5
	root.Left = &node1
	root.Right = &node2
	node1.Left = &node3
	node1.Right = &node4
	node2.Left = &node5
	node2.Right = &node6
	node4.Left = &node7
	node4.Right = &node8
	fmt.Println(lowestCommonAncestor(&root, &node1, &node5))
}
