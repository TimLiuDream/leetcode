package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

func main() {
	root := TreeNode{}
	root.Val = 1
	left := TreeNode{}
	left.Val = 2
	right := TreeNode{}
	right.Val = 3
	root.Left = &left
	root.Right = &right
	fmt.Println(isValidBST(&root))

	// root := TreeNode{}
	// root.Val = 5
	// left := TreeNode{}
	// left.Val = 1
	// right1 := TreeNode{}
	// right2 := TreeNode{}
	// right3 := TreeNode{}
	// right1.Val = 4
	// right2.Val = 3
	// right3.Val = 6
	// root.Left = &left
	// root.Right = &right1
	// right1.Left = &right2
	// right1.Right = &right3
	// fmt.Println(isValidBST(&root))
}
