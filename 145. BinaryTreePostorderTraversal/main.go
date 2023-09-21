package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		result := make([]int, 0)
		if node == nil {
			return result
		}
		result = append(result, traversal(node.Left)...)
		result = append(result, traversal(node.Right)...)
		result = append(result, node.Val)
		return result
	}
	return traversal(root)
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: nil,
		},
	}
	fmt.Println(postorderTraversal(root))

	fmt.Println(postorderTraversal(nil))

	root1 := &TreeNode{Val: 1}
	fmt.Println(postorderTraversal(root1))
}
