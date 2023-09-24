package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		pathP := path
		pathP += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, pathP)
		} else {
			pathP += "->"
			dfs(node.Left, pathP)
			dfs(node.Right, pathP)
		}
	}
	dfs(root, "")
	return paths
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(binaryTreePaths(root))

	root1 := &TreeNode{Val: 1}
	fmt.Println(binaryTreePaths(root1))
}
