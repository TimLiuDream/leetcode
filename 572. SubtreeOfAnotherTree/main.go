package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	sub := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSubtree(root, sub))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	var isSame func(node, subNode *TreeNode) bool
	isSame = func(node, subNode *TreeNode) bool {
		if node == nil && subNode == nil {
			return true
		}
		if node == nil || subNode == nil {
			return false
		}
		if node.Val == subNode.Val {
			return isSame(node.Left, subNode.Left) &&
				isSame(node.Right, subNode.Right)
		}
		return false

	}
	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
