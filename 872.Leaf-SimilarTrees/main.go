package main

import "fmt"

// Label 树
// Label 深度优先搜索

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := new(TreeNode)
	root1.Val = 1

	root1Left := new(TreeNode)
	root1Left.Val = 2

	root1Right := new(TreeNode)
	root1Right.Val = 3
	root1.Left = root1Left
	root1.Right = root1Right

	root2 := new(TreeNode)
	root2.Val = 1

	root2Left := new(TreeNode)
	root2Left.Val = 3

	root2Right := new(TreeNode)
	root2Right.Val = 2
	root2.Left = root2Left
	root2.Right = root2Right

	fmt.Println(leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		arr := make([]int, 0)
		if node.Left == nil && node.Right == nil {
			arr = append(arr, node.Val)
		} else {
			if node.Left != nil {
				arr = append(arr, dfs(node.Left)...)
			}
			if node.Right != nil {
				arr = append(arr, dfs(node.Right)...)
			}
		}
		return arr
	}

	root1Arr := dfs(root1)
	root2Arr := dfs(root2)
	if len(root1Arr) != len(root2Arr) {
		return false
	}
	for i, v := range root1Arr {
		if v != root2Arr[i] {
			return false
		}
	}
	return true
}
