package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历： 中 - 左 - 右
// 递归写法
func preorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		result := make([]int, 0)
		if node == nil {
			return result
		}
		result = append(result, node.Val)
		result = append(result, traversal(node.Left)...)
		result = append(result, traversal(node.Right)...)
		return result
	}
	return traversal(root)
}

// 迭代写法
func preorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			result = append(result, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return result
}

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
	res := preorderTraversal1(root)
	fmt.Println(res)

	res1 := preorderTraversal1(nil)
	fmt.Println(res1)

	root2 := &TreeNode{
		Val: 1,
	}
	res2 := preorderTraversal1(root2)
	fmt.Println(res2)
}
