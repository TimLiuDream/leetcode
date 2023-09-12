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
	res := inorderTraversal1(root)
	fmt.Println(res)

	res1 := inorderTraversal1(nil)
	fmt.Println(res1)

	root2 := &TreeNode{
		Val: 1,
	}
	res2 := inorderTraversal1(root2)
	fmt.Println(res2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度遍历，实用递归
// 中序遍历：左-中-右
func inorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		result := make([]int, 0)
		if node == nil {
			return result
		}
		result = append(result, traversal(node.Left)...)
		result = append(result, node.Val)
		result = append(result, traversal(node.Right)...)
		return result
	}
	return traversal(root)
}

// 迭代版本
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)

		curr = curr.Right
	}
	return res
}
