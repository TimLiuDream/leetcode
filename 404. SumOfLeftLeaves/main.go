package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var travesal func(node *TreeNode) int
	travesal = func(node *TreeNode) (ans int) {
		if node.Left != nil {
			if isLeaf(node.Left) {
				ans += node.Left.Val
			} else {
				ans += travesal(node.Left)
			}
		}
		if node.Right != nil && !isLeaf(node.Right) {
			ans += travesal(node.Right)
		}
		return
	}
	return travesal(root)
}

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(sumOfLeftLeaves(root))

	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(sumOfLeftLeaves(root1))
}
