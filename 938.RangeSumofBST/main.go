package main

import "fmt"

// Label 树
// Label 递归
// Label 深度优先搜索
func main() {
	// root = [10,5,15,3,7,null,18], low = 7, high = 15
	root := new(TreeNode)
	root.Val = 10

	node1 := new(TreeNode)
	node1.Val = 5

	node2 := new(TreeNode)
	node2.Val = 15

	node3 := new(TreeNode)
	node3.Val = 3

	node4 := new(TreeNode)
	node4.Val = 7

	node5 := new(TreeNode)
	node5.Val = 18

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Right = node5

	fmt.Println(rangeSumBST(root, 7, 15))

	//root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
	root = new(TreeNode)
	root.Val = 10

	node1 = new(TreeNode)
	node1.Val = 5

	node2 = new(TreeNode)
	node2.Val = 15

	node3 = new(TreeNode)
	node3.Val = 3

	node4 = new(TreeNode)
	node4.Val = 7

	node5 = new(TreeNode)
	node5.Val = 13

	node6 := new(TreeNode)
	node6.Val = 18

	node7 := new(TreeNode)
	node7.Val = 1

	node8 := new(TreeNode)
	node8.Val = 6

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node3.Left = node7
	node1.Right = node4
	node4.Left = node8
	node2.Left = node5
	node2.Right = node6

	fmt.Println(rangeSumBST(root, 6, 10))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	result := 0
	if root == nil {
		return result
	}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		value := 0
		if node == nil {
			return 0
		}
		if node.Left != nil {
			value += dfs(node.Left)
		}
		if node.Right != nil {
			value += dfs(node.Right)
		}
		if node.Val >= low && node.Val <= high {
			value += node.Val
			return value
		}
		return value
	}

	result += dfs(root)

	return result
}

// 官方做法
func rangeSumBST1(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
