package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(root))

	root = nil
	fmt.Println(inorderTraversal(root))

	root = &TreeNode{
		Val: 1,
	}
	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		part := make([]int, 0)
		if node == nil {
			return part
		}
		part = append(part, traversal(node.Left)...)
		part = append(part, node.Val)
		part = append(part, traversal(node.Right)...)
		return part
	}
	return traversal(root)
}
